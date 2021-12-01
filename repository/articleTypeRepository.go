package repository

import (
	"go-micro/database"
	"go-micro/entity"
)

type IArticleType interface {
	Get(id uint, fetchs []string) (*entity.ArticleType, error)
	GetAll(fetchs []string) ([]entity.ArticleType, error)
	Add(a *entity.ArticleType) (*entity.ArticleType, error)
	Delete(id uint) (*entity.ArticleType, error)
}

type ArticleType struct {
	DataBase *database.Database
}

func (this ArticleType) Get(id uint, fetchs []string) (*entity.ArticleType, error) {
	articleType := entity.ArticleType{}

	db := this.DataBase.GetDB()

	for _, fetch := range fetchs {
		db = db.Preload(fetch)
	}

	query := db.Find(&articleType, id)
	if query.Error != nil {
		return nil, query.Error
	} else if query.RowsAffected == 1 {
		return &articleType, nil
	} else {
		return nil, nil
	}
}

func (this ArticleType) GetAll(fetchs []string) ([]entity.ArticleType, error) {
	articleTypes := []entity.ArticleType{}

	db := this.DataBase.GetDB()

	for _, fetch := range fetchs {
		db = db.Preload(fetch)
	}

	query := db.Find(&articleTypes)
	return articleTypes, query.Error
}

func (this ArticleType) Add(at *entity.ArticleType) (*entity.ArticleType, error) {
	query := this.DataBase.GetDB().Create(at)
	if query.Error != nil {
		return nil, query.Error
	} else {
		return at, nil
	}
}

func (this ArticleType) Delete(id uint) (*entity.ArticleType, error) {
	articleType := entity.ArticleType{}
	query := this.DataBase.GetDB().Delete(&articleType, id)
	if query.Error != nil {
		return nil, query.Error
	} else {
		return &articleType, nil
	}
}
