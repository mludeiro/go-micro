package repository

import (
	"errors"
	"go-micro/database"
	"go-micro/entity"
)

func GetInvoice(id uint, fetchs []string) *entity.Invoice {
	dto := entity.Invoice{}
	db := database.GetDB()

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

func GetInvoices(fetchs []string) []entity.Invoice {
	dtos := []entity.Invoice{}
	db := database.GetDB()

	for _, fetch := range fetchs {
		db = db.Preload(fetch)
	}

	db.Find(&dtos)
	return dtos
}

func AddInvoice(a *entity.Invoice) (*entity.Invoice, error) {
	if database.GetDB().Create(a).RowsAffected != 1 {
		return nil, errors.New("Error creating new Invoice")
	}
	return a, nil
}
