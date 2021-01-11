package user

import (
	"kush-graphql/app/models"
)

type UserService interface {
	CreateUser(user *models.User) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	DeleteUser(id string) error
	GetUserByID(id string) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
	GetUserByEmail(emailAddress string) (*models.User, error)
}
