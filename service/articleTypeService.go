package service

import (
	"go-micro/entity"
	"go-micro/repository"
)

// Im limiting the add/delete function on this entity, simply by not exposing them on the interface
type IArticleType interface {
	Get(id uint, fetchs []string) *entity.ArticleType
	GetAll(fetchs []string) []entity.ArticleType
	// Add(a *entity.ArticleType) (*entity.ArticleType, error)
	// Delete(id uint) *entity.ArticleType
}

// bypass the get/getAll
type ArticleType struct {
	repository.IArticleType
}
