package repository

import (
	"go-micro/database"

	"gorm.io/gorm"
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

func (repo *Repository[T]) getQueryDB(query *Query) *gorm.DB {
	tx := repo.DataBase.GetDB()

	for _, fetch := range query.Fetchs {
		tx = tx.Preload(fetch)
	}

	for _, cond := range query.Conditions {
		switch cond.Comparator {
		case "eq":
			tx = tx.Where(cond.Field, cond.Value)
		case "lk":
			tx = tx.Where(cond.Field, cond.Value)
		}
	}

	for _, order := range query.OrderBy {
		tx = tx.Order(order)
	}

	return tx
}

func getResult[T any](tx gorm.DB, query *Query) *ResultSet[T] {
	var rset ResultSet[T]
	dataQuery := tx.Limit(int(query.PageSize)).Offset(int(query.PageNumber * query.PageSize)).Find(&rset.Data)

	rset.Error = dataQuery.Error
	rset.Page = query.PageNumber
	tx.Model(rset.Data).Count(&rset.Total)

	return &rset
}

func (this *Repository[T]) GetAll(query *Query) *ResultSet[T] {
	var rset ResultSet[T]

	dataQuery := this.getQueryDB(query)
	dataQuery = dataQuery.Limit(int(query.PageSize)).Offset(int(query.PageNumber * query.PageSize)).Find(&rset.Data)

	rset.Error = dataQuery.Error
	rset.Page = query.PageNumber
	dataQuery.Model(rset.Data).Count(&rset.Total)

	return &rset
}

func (this *Repository[T]) Add(a *T) (*T, error) {
	query := this.DataBase.GetDB().Create(a)
	if query.Error != nil {
		return nil, query.Error
	}
	return a, nil
}

func (this *Repository[T]) Delete(id uint) (*T, error) {
	var dto T
	query := this.DataBase.GetDB().Where("deleted_at is NULL").Delete(&dto, id)
	if query.Error != nil {
		return nil, query.Error
	} else {
		if query.RowsAffected == 1 {
			return &dto, nil
		} else {
			return nil, nil
		}
	}
}
