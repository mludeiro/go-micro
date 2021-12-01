package service

import (
	"go-micro/entity"
	"go-micro/repository"
)

type IActicle interface {
	Get(uint, []string) (*entity.Article, error)
	GetAll([]string) ([]entity.Article, error)
	Add(*entity.Article) (*entity.Article, error)
	Delete(uint) (*entity.Article, error)
}

type Article struct {
	repository.IArticle
}

// if you want to, you can wrap or redefine the repository method
func (a Article) GetAll(fetchs []string) ([]entity.Article, error) {
	return a.IArticle.GetAll(fetchs)
}
