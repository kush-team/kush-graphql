package persistence

import (
	"kush-graphql/app/domain/repository/article"
	"kush-graphql/app/models"

	"github.com/jinzhu/gorm"
)

type articleService struct {
	db *gorm.DB
}

func NewArticle(db *gorm.DB) *articleService {
	return &articleService{
		db,
	}
}

var _ article.ArticleService = &articleService{}

func (s *articleService) CreateArticle(article *models.Article) (*models.Article, error) {

	err := s.db.Create(&article).Error
	if err != nil {
		return nil, err
	}

	return article, nil
}

func (s *articleService) GetAllArticles() ([]*models.Article, error) {

	var articles []*models.Article

	err := s.db.Preload("Category").Preload("Author").Find(&articles).Error
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (s *articleService) GetArticleByID(id string) (*models.Article, error) {

	article := &models.Article{}
	err := s.db.Preload("Author").Preload("Category").Where("id = ?", id).Take(&article).Error
	if err != nil {
		return nil, err
	}

	return article, nil
}
