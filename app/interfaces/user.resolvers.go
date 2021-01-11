package interfaces

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"kush-graphql/app/auth"
	"kush-graphql/app/models"
	"kush-graphql/helpers"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (r *mutationResolver) CreateUser(ctx context.Context, user models.UserInput) (*models.UserResponse, error) {
	hashedPassword, err := HashPassword(user.Password)

	usr := &models.User{
		Username:     user.Username,
		Password:     hashedPassword,
		EmailAddress: user.EmailAddress,
		Role:         user.Role,
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
	usr, err := r.UserService.GetUserByID(id)
	if err != nil {
		log.Println("Error getting the user to update: ", err)
		return &models.UserResponse{
			Message: "Error getting the user",
			Status:  http.StatusUnprocessableEntity,
		}, nil
	}

	usr.Username = user.Username

	if ok, errorString := helpers.ValidateInputs(*usr); !ok {
		return &models.UserResponse{
			Message: errorString,
			Status:  http.StatusUnprocessableEntity,
		}, nil
	}

	userCreated, err := r.UserService.UpdateUser(usr)
	if err != nil {
		log.Println("User updating error: ", err)
		return &models.UserResponse{
			Message: "Error updating User",
			Status:  http.StatusInternalServerError,
		}, nil
	}

	return &models.UserResponse{
		Message: "Successfully updated user",
		Status:  http.StatusOK,
		Data:    userCreated,
	}, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*models.UserResponse, error) {
	err := r.UserService.DeleteUser(id)
	if err != nil {
		return &models.UserResponse{
			Message: "Something went wrong deleting the User.",
			Status:  http.StatusInternalServerError,
		}, nil
	}

	return &models.UserResponse{
		Message: "Successfully deleted User",
		Status:  http.StatusOK,
	}, nil
}

func (r *mutationResolver) Login(ctx context.Context, emailAddress string, password string) (*models.LoginResponse, error) {
	user, err := r.UserService.GetUserByEmail(emailAddress)

	if err != nil {
		return &models.LoginResponse{
			Message: "Something went wrong login the User.",
			Status:  http.StatusInternalServerError,
		}, nil
	}

	match := CheckPasswordHash(password, user.Password)

	if match {
		token, err := auth.GenerateToken(emailAddress)
		if err == nil {
			return &models.LoginResponse{
				Message: "Successfully login",
				Status:  http.StatusOK,
				Token:   &token,
			}, nil
		}
	}

	return &models.LoginResponse{
		Message: "Something went wrong login the User.",
		Status:  http.StatusUnprocessableEntity,
	}, nil
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
