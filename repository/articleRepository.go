package repository

import (
	"errors"
	"go-micro/database"
	"go-micro/entity"
)

type IArticle interface {
	Get(id uint, fetchs []string) *entity.Article
	GetAll(fetchs []string) []entity.Article
	Add(a *entity.Article) (*entity.Article, error)
	Delete(id uint) *entity.Article
}

type Article struct {
	DataBase *database.Database
}

func (this *Article) Get(id uint, fetchs []string) *entity.Article {
	article := entity.Article{}
	db := this.DataBase.GetDB()

	for _, fetch := range fetchs {
		db = db.Preload(fetch)
	}

	rows := db.Find(&article, id).RowsAffected
	if rows == 1 {
		return &article
	} else {
		return nil
	}
}

func (this *Article) GetAll(fetchs []string) []entity.Article {
	articles := []entity.Article{}
	db := this.DataBase.GetDB()

	for _, fetch := range fetchs {
		db = db.Preload(fetch)
	}

	db.Find(&articles)
	return articles
}

func (this *Article) Add(a *entity.Article) (*entity.Article, error) {
	if this.DataBase.GetDB().Create(a).RowsAffected != 1 {
		return nil, errors.New("Error creating new article")
	}
	return a, nil
}

func (this *Article) Delete(id uint) *entity.Article {
	article := entity.Article{}
	rows := this.DataBase.GetDB().Where("deleted_at is NULL").Delete(&article, id).RowsAffected
	if rows == 1 {
		return &article
	} else {
		return nil
	}
}
