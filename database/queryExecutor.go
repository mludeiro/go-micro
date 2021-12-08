package database

import (
	"go-micro/entity"

	"gorm.io/gorm"
)

type QueryExecutor struct {
	trx        *gorm.DB
	pageSize   uint
	pageNumber uint
}

func (this *QueryExecutor) GetResult(set *entity.ResultSet, dest interface{}) error {

	dataQuery := this.trx.Limit(int(this.pageSize)).Offset(int(this.pageNumber * this.pageSize)).Find(dest)

	if dataQuery.Error != nil {
		return dataQuery.Error
	}

	set.Page = this.pageNumber
	this.trx.Model(dest).Count(&set.Total)

	return nil
}
