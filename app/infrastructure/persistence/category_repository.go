package persistence

import (
	"kush-graphql/app/domain/repository/category"
	"kush-graphql/app/models"

	"github.com/jinzhu/gorm"
)

type categoryService struct {
	db *gorm.DB
}

func NewCategory(db *gorm.DB) *categoryService {
	return &categoryService{
		db,
	}
}

var _ category.CategoryService = &categoryService{}

func (s *categoryService) CreateCategory(category *models.Category) (*models.Category, error) {

	err := s.db.Create(&category).Error
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s *categoryService) UpdateCategory(category *models.Category) (*models.Category, error) {

	err := s.db.Save(&category).Error
	if err != nil {
		return nil, err
	}

	return category, nil

}

func (s *categoryService) DeleteCategory(id string) error {

	cat := &models.Category{}

	err := s.db.Where("id = ?", id).Delete(cat).Error
	if err != nil {
		return err
	}

	return nil

}

func (s *categoryService) GetCategoryByID(id string) (*models.Category, error) {

	cat := &models.Category{}

	err := s.db.Where("id = ?", id).Take(&cat).Error
	if err != nil {
		return nil, err
	}

	return cat, nil
}

func (s *categoryService) GetAllCategorys() ([]*models.Category, error) {

	var categories []*models.Category

	err := s.db.Preload("Articles").Preload("Articles.Author").Preload("Articles.Category").Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return categories, nil
}
