package main

import (
	"context"
	"fmt"
	"kush-graphql/app/auth"
	"kush-graphql/app/domain/repository/article"
	"kush-graphql/app/domain/repository/category"
	"kush-graphql/app/domain/repository/theme"
	"kush-graphql/app/domain/repository/user"
	"kush-graphql/app/generated"
	"kush-graphql/app/infrastructure/db"
	"kush-graphql/app/infrastructure/persistence"
	"kush-graphql/app/interfaces"
	"kush-graphql/app/models"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	var (
		defaultPort      = "8080"
		databaseUser     = os.Getenv("DATABASE_USER")
		databaseName     = os.Getenv("DATABASE_NAME")
		databaseHost     = os.Getenv("DATABASE_HOST")
		databasePort     = os.Getenv("DATABASE_PORT")
		databasePassword = os.Getenv("DATABASE_PASSWORD")
	)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	router := chi.NewRouter()

	router.Use(auth.Middleware())

	dbConn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", databaseHost, databasePort, databaseUser, databaseName, databasePassword)

	conn := db.OpenDB(dbConn)

	var userService user.UserService
	var categoryService category.CategoryService
	var articleService article.ArticleService
	var themeService theme.ThemeService

	userService = persistence.NewUser(conn)
	categoryService = persistence.NewCategory(conn)
	articleService = persistence.NewArticle(conn)
	themeService = persistence.NewTheme(conn)

	c := generated.Config{Resolvers: &interfaces.Resolver{
		UserService:     userService,
		CategoryService: categoryService,
		ArticleService:  articleService,
		ThemeService:    themeService,
	}}

	c.Directives.HasRole = func(ctx context.Context, obj interface{}, next graphql.Resolver, role models.Role) (interface{}, error) {
		// TO-DO Get headers and validate token
		user := auth.ForContext(ctx)

		if user == nil {
			return nil, fmt.Errorf("Access denied")
		}

		user, err := userService.GetUserByEmail(user.EmailAddress)

		if err == nil && user != nil && user.Role != models.RoleAdmin {
			return nil, fmt.Errorf("Access denied")
		} else {
			return next(ctx)
		}
	}

	srv := handler.New(generated.NewExecutableSchema(c))

	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	srv.SetQueryCache(lru.New(1000))

	srv.Use(extension.Introspection{})

	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	router.Handle("/playground", playground.Handler("GraphQL playground", "/graphql"))
	router.Handle("/graphql", srv)

	FileServer(router, "/", http.Dir("./web/dist/kush-blog"))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
