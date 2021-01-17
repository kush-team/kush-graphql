package persistence

import (
	"kush-graphql/app/domain/repository/theme"
	"kush-graphql/app/models"

	"github.com/jinzhu/gorm"
)

type themeService struct {
	db *gorm.DB
}

func NewTheme(db *gorm.DB) *themeService {
	return &themeService{
		db,
	}
}

var _ theme.ThemeService = &themeService{}

func (s *themeService) CreateTheme(theme *models.Theme) (*models.Theme, error) {

	err := s.db.Create(&theme).Error
	if err != nil {
		return nil, err
	}

	return theme, nil
}

func (s *themeService) UpdateTheme(theme *models.Theme) (*models.Theme, error) {

	err := s.db.Save(&theme).Error
	if err != nil {
		return nil, err
	}

	return theme, nil

}

func (s *themeService) DeleteTheme(id string) error {

	theme := &models.Theme{}

	err := s.db.Where("id = ?", id).Delete(theme).Error
	if err != nil {
		return err
	}

	return nil

}

func (s *themeService) GetThemeByID(id string) (*models.Theme, error) {

	theme := &models.Theme{}

	err := s.db.Preload("Author").Where("id = ?", id).Take(&theme).Error
	if err != nil {
		return nil, err
	}

	return theme, nil
}

func (s *themeService) GetThemeByName(name *string) (*models.Theme, error) {

	theme := &models.Theme{}

	err := s.db.Preload("Author").Where("name = ?", &name).Take(&theme).Error
	if err != nil {
		return nil, err
	}

	return theme, nil
}

func (s *themeService) GetAllThemes() ([]*models.Theme, error) {

	var themes []*models.Theme

	err := s.db.Preload("Author").Find(&themes).Error

	if err != nil {
		return nil, err
	}

	return themes, nil
}
