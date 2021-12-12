package service

import (
	"go-micro/entity"
	"go-micro/repository"
)

type IActicleService interface {
	Get(uint, []string) (*entity.Article, error)
	GetAll(query entity.Query) (entity.ArticleResultSet, error)
	Add(*entity.Article) (*entity.Article, error)
	Delete(uint) (*entity.Article, error)
}

type Article struct {
	repository.IArticleRepository
}

// if you want to, you can wrap or redefine the repository method
func (a *Article) GetAll(query entity.Query) (entity.ArticleResultSet, error) {
	return a.IArticleRepository.GetAll(query)
}
