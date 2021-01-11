package interfaces

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"kush-graphql/app/generated"
	"kush-graphql/app/models"
)

func (r *mutationResolver) CreateArticle(ctx context.Context, article models.ArticleInput) (*models.ArticleResponse, error) {
	panic(fmt.Errorf("not implemented"))
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
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetArticlesByTags(ctx context.Context, tags []*models.TagInput) (*models.ArticleResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetArticlesByCategory(ctx context.Context, category models.CategoryInput) (*models.ArticleResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
