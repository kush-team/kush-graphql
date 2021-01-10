package interfaces

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"kush-graphql/app/models"
)

func (r *mutationResolver) CreateTag(ctx context.Context, tag models.TagInput) (*models.TagResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateTag(ctx context.Context, id string, tag models.TagInput) (*models.TagResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteTag(ctx context.Context, id string) (*models.TagResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetTagByID(ctx context.Context, id string) (*models.TagResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetAllTags(ctx context.Context) (*models.TagResponse, error) {
	panic(fmt.Errorf("not implemented"))
}
