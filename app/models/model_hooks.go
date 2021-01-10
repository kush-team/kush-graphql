package models

import (
	"github.com/jinzhu/gorm"
	"github.com/twinj/uuid"
)

func (mod *User) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("id", uuid.String())
}


func (mod *Tag) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("id", uuid.String())
}

func (mod *Article) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("id", uuid.String())
}

func (mod *Category) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("id", uuid.String())
}


