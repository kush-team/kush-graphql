package persistence

import (
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
