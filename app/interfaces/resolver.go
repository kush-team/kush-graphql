package interfaces

import (
	"kush-graphql/app/domain/repository/article"
	"kush-graphql/app/domain/repository/category"
	"kush-graphql/app/domain/repository/theme"
	"kush-graphql/app/domain/repository/user"
)

// Resolver This file will not be regenerated automatically.
type Resolver struct {
	UserService     user.UserService
	CategoryService category.CategoryService
	ArticleService  article.ArticleService
	ThemeService    theme.ThemeService
}
