package article

import (
	"kush-graphql/app/models"
)

type ArticleService interface {
	CreateArticle(Article *models.Article) (*models.Article, error)
	//UpdateArticle(Article *models.Article) (*models.Article, error)
	//DeleteArticle(id string) error
	//GetArticleByID(id string) (*models.Article, error)
	GetAllArticles() ([]*models.Article, error)
	//GetArticlesByTags(tags *[]models.Tag) ([]*models.Article, error)
	//GetArticlesByCategory(category *models.Category) ([]*models.Article, error)
}
