package main

import (
	"context"
	"fmt"
	"kush-graphql/app/auth"
	"kush-graphql/app/domain/repository/category"
	"kush-graphql/app/domain/repository/user"
	"kush-graphql/app/generated"
	"kush-graphql/app/infrastructure/db"
	"kush-graphql/app/infrastructure/persistence"
	"kush-graphql/app/interfaces"
	"kush-graphql/app/models"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
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

	userService = persistence.NewUser(conn)
	categoryService = persistence.NewCategory(conn)

	c := generated.Config{Resolvers: &interfaces.Resolver{
		UserService:     userService,
		CategoryService: categoryService,
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

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
