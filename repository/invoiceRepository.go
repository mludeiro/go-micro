package repository

import (
	"errors"
	"go-micro/database"
	"go-micro/entity"
)

type IInvoiceRepository interface {
	Get(id uint, fetchs []string) (*entity.Invoice, error)
	GetAll(fetchs []string) ([]entity.Invoice, error)
	Add(a *entity.Invoice) (*entity.Invoice, error)
}

type Invoice struct {
	DataBase *database.Database
}

func (this *Invoice) Get(id uint, fetchs []string) (*entity.Invoice, error) {
	dto := entity.Invoice{}

	db := this.DataBase.GetDB()

	for _, fetch := range fetchs {
		db = db.Preload(fetch)
	}

	query := db.Find(&dto, id)
	if query.Error != nil {
		return nil, query.Error
	} else if query.RowsAffected == 1 {
		return &dto, nil
	} else {
		return nil, nil
	}
}

func (this *Invoice) GetAll(fetchs []string) ([]entity.Invoice, error) {
	dto := []entity.Invoice{}

	db := this.DataBase.GetDB()

	for _, fetch := range fetchs {
		db = db.Preload(fetch)
	}

	query := db.Find(&dto)
	return dto, query.Error
}

func (this *Invoice) Add(a *entity.Invoice) (*entity.Invoice, error) {
	if this.DataBase.GetDB().Create(a).RowsAffected != 1 {
		return nil, errors.New("Error creating new Invoice")
	}
	return a, nil
}
