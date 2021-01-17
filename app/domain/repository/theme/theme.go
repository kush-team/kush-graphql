package theme

import (
	"kush-graphql/app/models"
)

type ThemeService interface {
	CreateTheme(Theme *models.Theme) (*models.Theme, error)
	UpdateTheme(Theme *models.Theme) (*models.Theme, error)
	DeleteTheme(id string) error
	GetThemeByID(id string) (*models.Theme, error)
	GetThemeByName(name *string) (*models.Theme, error)
	GetAllThemes() ([]*models.Theme, error)
}
