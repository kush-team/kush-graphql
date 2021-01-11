package interfaces

import (
	"kush-graphql/app/domain/repository/category"
	"kush-graphql/app/domain/repository/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserService     user.UserService
	CategoryService category.CategoryService
}
