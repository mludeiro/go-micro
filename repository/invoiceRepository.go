package repository

import (
	"errors"
	"go-micro/database"
	"go-micro/entity"
)

type IInvoice interface {
	Get(id uint, fetchs []string) *entity.Invoice
	GetAll(fetchs []string) []entity.Invoice
	Add(a *entity.Invoice) (*entity.Invoice, error)
}

type Invoice struct {
	DataBase database.Database
}

func (this *Invoice) Get(id uint, fetchs []string) *entity.Invoice {
	dto := entity.Invoice{}
	db := this.DataBase.GetDB()

	for _, fetch := range fetchs {
		db = db.Preload(fetch)
	}

	rows := db.Find(&dto, id).RowsAffected
	if rows == 1 {
		return &dto
	} else {
		return nil
	}
}

func (this *Invoice) GetAll(fetchs []string) []entity.Invoice {
	dtos := []entity.Invoice{}
	db := this.DataBase.GetDB()

	for _, fetch := range fetchs {
		db = db.Preload(fetch)
	}

	db.Find(&dtos)
	return dtos
}

func (this *Invoice) Add(a *entity.Invoice) (*entity.Invoice, error) {
	if this.DataBase.GetDB().Create(a).RowsAffected != 1 {
		return nil, errors.New("Error creating new Invoice")
	}
	return a, nil
}
