package repository

import (
	"go-micro/database"
	"go-micro/entity"
)

type IArticleRepository interface {
	Get(id uint, fetchs []string) (*entity.Article, error)
	GetAll(query database.Query) ([]entity.Article, error)
	Add(a *entity.Article) (*entity.Article, error)
	Delete(id uint) (*entity.Article, error)
}

type Article struct {
	DataBase *database.Database
}

func (this Article) Get(id uint, fetchs []string) (*entity.Article, error) {
	article := entity.Article{}
	db := this.DataBase.GetDB()

	for _, fetch := range fetchs {
		db = db.Preload(fetch)
	}

	query := db.Find(&article, id)
	if query.Error == nil && query.RowsAffected == 1 {
		return &article, nil
	} else {
		return nil, query.Error
	}
}

func (this Article) GetAll(query database.Query) ([]entity.Article, error) {
	articles := []entity.Article{}
	db := this.DataBase.GetQueryDB(query)

	err := db.Find(&articles).Error
	return articles, err
}

func (this Article) Add(a *entity.Article) (*entity.Article, error) {
	query := this.DataBase.GetDB().Create(a)
	if query.Error != nil {
		return nil, query.Error
	}
	return a, nil
}

func (this Article) Delete(id uint) (*entity.Article, error) {
	article := entity.Article{}
	query := this.DataBase.GetDB().Where("deleted_at is NULL").Delete(&article, id)
	if query.Error != nil {
		return nil, query.Error
	} else {
		if query.RowsAffected == 1 {
			return &article, nil
		} else {
			return nil, nil
		}
	}
}
