package interfaces

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"kush-graphql/app/models"
	"kush-graphql/helpers"
	"log"
	"net/http"
)

func (r *mutationResolver) CreateCategory(ctx context.Context, category models.CategoryInput) (*models.CategoryResponse, error) {
	cat := &models.Category{
		Name: category.Name,
	}

	if ok, errorString := helpers.ValidateInputs(*cat); !ok {
		return &models.CategoryResponse{
			Message: errorString,
			Status:  http.StatusUnprocessableEntity,
		}, nil
	}

	catCreated, err := r.CategoryService.CreateCategory(cat)

	if err != nil {
		log.Println("Category creation error: ", err)
		return &models.CategoryResponse{
			Message: "Error creating Category",
			Status:  http.StatusInternalServerError,
		}, nil
	}

	return &models.CategoryResponse{
		Message: "Successfully created Category",
		Status:  http.StatusCreated,
		Data:    catCreated,
	}, nil
}

func (r *mutationResolver) UpdateCategory(ctx context.Context, id string, category models.CategoryInput) (*models.CategoryResponse, error) {
	cat, err := r.CategoryService.GetCategoryByID(id)
	if err != nil {
		log.Println("Error getting the category to update: ", err)
		return &models.CategoryResponse{
			Message: "Error getting the category",
			Status:  http.StatusUnprocessableEntity,
		}, nil
	}

	cat.Name = category.Name

	if ok, errorString := helpers.ValidateInputs(*cat); !ok {
		return &models.CategoryResponse{
			Message: errorString,
			Status:  http.StatusUnprocessableEntity,
		}, nil
	}

	categoryUpdated, err := r.CategoryService.UpdateCategory(cat)
	if err != nil {
		log.Println("Category updating error: ", err)
		return &models.CategoryResponse{
			Message: "Error updating Category",
			Status:  http.StatusInternalServerError,
		}, nil
	}

	return &models.CategoryResponse{
		Message: "Successfully updated Category",
		Status:  http.StatusOK,
		Data:    categoryUpdated,
	}, nil
}

func (r *mutationResolver) DeleteCategory(ctx context.Context, id string) (*models.CategoryResponse, error) {
	err := r.CategoryService.DeleteCategory(id)
	if err != nil {
		return &models.CategoryResponse{
			Message: "Something went wrong deleting the Category.",
			Status:  http.StatusInternalServerError,
		}, nil
	}

	return &models.CategoryResponse{
		Message: "Successfully deleted Category",
		Status:  http.StatusOK,
	}, nil
}

func (r *queryResolver) GetCategoryByID(ctx context.Context, id string) (*models.CategoryResponse, error) {
	category, err := r.CategoryService.GetCategoryByID(id)
	if err != nil {
		log.Println("getting category error: ", err)
		return &models.CategoryResponse{
			Message: "Something went wrong getting the category.",
			Status:  http.StatusInternalServerError,
		}, nil
	}

	return &models.CategoryResponse{
		Message: "Successfully retrieved category",
		Status:  http.StatusOK,
		Data:    category,
	}, nil
}

func (r *queryResolver) GetAllCategorys(ctx context.Context) (*models.CategoryResponse, error) {
	categories, err := r.CategoryService.GetAllCategorys()
	if err != nil {
		log.Println("getting all categories error: ", err)
		return &models.CategoryResponse{
			Message: "Something went wrong getting all categories.",
			Status:  http.StatusInternalServerError,
		}, nil
	}

	return &models.CategoryResponse{
		Message:  "Successfully retrieved all categories",
		Status:   http.StatusOK,
		DataList: categories,
	}, nil
}
