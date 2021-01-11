package interfaces

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"kush-graphql/app/generated"
	"kush-graphql/app/models"
	"kush-graphql/helpers"
	"log"
	"net/http"
	"time"

	"github.com/markbates/going/randx"
)

func (r *mutationResolver) CreateArticle(ctx context.Context, article models.ArticleInput) (*models.ArticleResponse, error) {
	art := &models.Article{
		Name:       article.Name,
		AuthorID:   article.AuthorID,
		Brief:      article.Brief,
		Content:    article.Content,
		CategoryID: article.CategoryID,
	}

	art.CreatedAt = time.Now()
	art.UpdatedAt = time.Now()

	if ok, errorString := helpers.ValidateInputs(*art); !ok {
		return &models.ArticleResponse{
			Message: errorString,
			Status:  http.StatusUnprocessableEntity,
		}, nil
	}

	artCreated, err := r.ArticleService.CreateArticle(art)

	if err != nil {
		log.Println("Article creation error: ", err)
		return &models.ArticleResponse{
			Message: "Error creating Article",
			Status:  http.StatusInternalServerError,
		}, nil
	}

	for _, observer := range articlePublishedChannel {
		observer <- artCreated
	}

	return &models.ArticleResponse{
		Message: "Successfully created Article",
		Status:  http.StatusCreated,
		Data:    artCreated,
	}, nil
}

func (r *mutationResolver) UpdateArticle(ctx context.Context, id string, article models.ArticleInput) (*models.ArticleResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteArticle(ctx context.Context, id string) (*models.ArticleResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetArticleByID(ctx context.Context, id string) (*models.ArticleResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetAllArticles(ctx context.Context) (*models.ArticleResponse, error) {
	articles, err := r.ArticleService.GetAllArticles()
	if err != nil {
		log.Println("getting all articles error: ", err)
		return &models.ArticleResponse{
			Message: "Something went wrong getting all articles.",
			Status:  http.StatusInternalServerError,
		}, nil
	}

	return &models.ArticleResponse{
		Message:  "Successfully retrieved all articles",
		Status:   http.StatusOK,
		DataList: articles,
	}, nil
}

func (r *queryResolver) GetArticlesByTags(ctx context.Context, tags []*models.TagInput) (*models.ArticleResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetArticlesByCategory(ctx context.Context, category models.CategoryInput) (*models.ArticleResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) ArticleAdded(ctx context.Context) (<-chan *models.Article, error) {
	id := randx.String(8)

	articleEvent := make(chan *models.Article, 1)
	go func() {
		<-ctx.Done()
	}()

	articlePublishedChannel[id] = articleEvent

	return articleEvent, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var articlePublishedChannel map[string]chan *models.Article

func init() {
	articlePublishedChannel = map[string]chan *models.Article{}
}
