package interfaces

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"kush-graphql/app/models"
)

func (r *mutationResolver) CreateCategory(ctx context.Context, category models.CategoryInput) (*models.CategoryResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateCategory(ctx context.Context, id string, category models.CategoryInput) (*models.CategoryResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteCategory(ctx context.Context, id string) (*models.CategoryResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetCategoryByID(ctx context.Context, id string) (*models.CategoryResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetAllCategorys(ctx context.Context) (*models.CategoryResponse, error) {
	panic(fmt.Errorf("not implemented"))
}
