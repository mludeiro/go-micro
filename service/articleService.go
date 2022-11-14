package service

import (
	"go-micro/entity"
	"go-micro/repository"
)

type IActicleService interface {
	Get(uint, []string) (*entity.Article, error)
	GetAll(query *repository.Query) (*repository.ResultSet[entity.Article], error)
	Add(*entity.Article) (*entity.Article, error)
	Delete(uint) (*entity.Article, error)
}

type Article struct {
	Repo *repository.Repository[entity.Article]
}

func (a *Article) Get(id uint, fetchs []string) (*entity.Article, error) {
	return a.Repo.Get(id, fetchs)
}

func (a *Article) GetAll(query *repository.Query) (*repository.ResultSet[entity.Article], error) {
	return a.Repo.GetAll(query), nil
}

func (a *Article) Add(dto *entity.Article) (*entity.Article, error) {
	return a.Repo.Add(dto)
}

func (a *Article) Delete(id uint) (*entity.Article, error) {
	return a.Repo.Delete(id)
}
