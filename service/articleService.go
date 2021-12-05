package service

import (
	"go-micro/database"
	"go-micro/entity"
	"go-micro/repository"
)

type IActicleService interface {
	Get(uint, []string) (*entity.Article, error)
	GetAll(query database.Query) ([]entity.Article, error)
	Add(*entity.Article) (*entity.Article, error)
	Delete(uint) (*entity.Article, error)
}

type Article struct {
	repository.IArticleRepository
}

// if you want to, you can wrap or redefine the repository method
func (a Article) GetAll(query database.Query) ([]entity.Article, error) {
	return a.IArticleRepository.GetAll(query)
}
