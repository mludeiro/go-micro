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
	Repository repository.IArticle
}

func (this Article) Get(id uint, fetchs []string) (*entity.Article, error) {
	return this.Repository.Get(id, fetchs)
}

func (this Article) GetAll(fetchs []string) ([]entity.Article, error) {
	return this.Repository.GetAll(fetchs)
}

func (this Article) Add(dto *entity.Article) (*entity.Article, error) {
	return this.Repository.Add(dto)
}

func (this Article) Delete(id uint) (*entity.Article, error) {
	return this.Repository.Delete(id)
}
