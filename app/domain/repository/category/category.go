package category

import (
	"kush-graphql/app/models"
)

type CategoryService interface {
	CreateCategory(Category *models.Category) (*models.Category, error)
	UpdateCategory(Category *models.Category) (*models.Category, error)
	DeleteCategory(id string) error
	GetCategoryByID(id string) (*models.Category, error)
	GetAllCategorys() ([]*models.Category, error)
}
