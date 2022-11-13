package repository

import (
	"go-micro/database"
	"go-micro/entity"
)

// type IArticleRepository interface {
// 	Get(id uint, fetchs []string) (*entity.Article, error)
// 	GetAll(query entity.Query) (entity.ArticleResultSet, error)
// 	Add(a *entity.Article) (*entity.Article, error)
// 	Delete(id uint) (*entity.Article, error)
// }

type Repository[T any] struct {
	DataBase *database.Database
}

func (this *Repository[T]) Get(id uint, fetchs []string) (*T, error) {
	var dto T
	db := this.DataBase.GetDB()

	for _, fetch := range fetchs {
		db = db.Preload(fetch)
	}

	query := db.Find(&dto, id)
	if query.Error == nil && query.RowsAffected == 1 {
		return &dto, nil
	} else {
		return nil, query.Error
	}
}

func (this *Repository[T]) GetAll(query entity.Query) (entity.ResultSet[T], error) {
	dtos := entity.ResultSet[T]{Query: query}
	P
	dataQuery := this.DataBase.GetDB().Limit(int(query.PageSize)).Offset(int(query.PageNumber * query.PageSize)).Find(dest)

	if dataQuery.Error != nil {
		return dataQuery.Error
	}

	dtos.Page = query.PageNumber
	this.trx.Model(dest).Count(&set.Total)

	err := this.DataBase.GetQueryDB(query).GetResult(&dtos, &dtos.Data)

	return articles, err
}

func (this *Repository[T]) Add(a *entity.Article) (*entity.Article, error) {
	query := this.DataBase.GetDB().Create(a)
	if query.Error != nil {
		return nil, query.Error
	}
	return a, nil
}

func (this *Repository[T]) Delete(id uint) (*entity.Article, error) {
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
