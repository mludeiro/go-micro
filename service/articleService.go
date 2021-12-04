package service

import (
	"go-micro/entity"
	"go-micro/repository"
)

type IActicleService interface {
	Get(uint, []string) (*entity.Article, error)
	GetAll([]string) ([]entity.Article, error)
	Add(*entity.Article) (*entity.Article, error)
	Delete(uint) (*entity.Article, error)
}

type Article struct {
	repository.IArticleRepository
}

// if you want to, you can wrap or redefine the repository method
func (a Article) GetAll(fetchs []string) ([]entity.Article, error) {
	return a.IArticleRepository.GetAll(fetchs)
}
