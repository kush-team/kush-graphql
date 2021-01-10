package interfaces

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"kush-graphql/app/models"
	"kush-graphql/helpers"
	"log"
	"net/http"
)

func (r *mutationResolver) CreateUser(ctx context.Context, user models.UserInput) (*models.UserResponse, error) {
	usr := &models.User{
		Username:     user.Username,
		Password:     user.Password,
		EmailAddress: user.EmailAddress,
	}

	if ok, errorString := helpers.ValidateInputs(*usr); !ok {
		return &models.UserResponse{
			Message: errorString,
			Status:  http.StatusUnprocessableEntity,
		}, nil
	}

	userCreated, err := r.UserService.CreateUser(usr)

	if err != nil {
		log.Println("User creation error: ", err)
		return &models.UserResponse{
			Message: "Error creating user",
			Status:  http.StatusInternalServerError,
		}, nil
	}

	return &models.UserResponse{
		Message: "Successfully created user",
		Status:  http.StatusCreated,
		Data:    userCreated,
	}, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, user models.UserInput) (*models.UserResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*models.UserResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetUserByID(ctx context.Context, id string) (*models.UserResponse, error) {
	user, err := r.UserService.GetUserByID(id)
	if err != nil {
		log.Println("getting user error: ", err)
		return &models.UserResponse{
			Message: "Something went wrong getting the user.",
			Status:  http.StatusInternalServerError,
		}, nil
	}

	return &models.UserResponse{
		Message: "Successfully retrieved user",
		Status:  http.StatusOK,
		Data:    user,
	}, nil
}

func (r *queryResolver) GetAllUsers(ctx context.Context) (*models.UserResponse, error) {
	users, err := r.UserService.GetAllUsers()
	if err != nil {
		log.Println("getting all users error: ", err)
		return &models.UserResponse{
			Message: "Something went wrong getting all users.",
			Status:  http.StatusInternalServerError,
		}, nil
	}

	return &models.UserResponse{
		Message:  "Successfully retrieved all users",
		Status:   http.StatusOK,
		DataList: users,
	}, nil
}
