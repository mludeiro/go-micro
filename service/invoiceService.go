package service

import (
	"go-micro/entity"
	"go-micro/repository"
)

type IInvoice interface {
	Get(uint, []string) *entity.Invoice
	GetAll([]string) []entity.Invoice
	Add(*entity.Invoice) (*entity.Invoice, error)
}

type Invoice struct {
	Repository repository.IInvoice
}

func (this Invoice) Get(id uint, fetchs []string) *entity.Invoice {
	return this.Repository.Get(id, fetchs)
}

func (this Invoice) GetAll(fetchs []string) []entity.Invoice {
	return this.Repository.GetAll(fetchs)
}

func (this Invoice) Add(dto *entity.Invoice) (*entity.Invoice, error) {
	return this.Repository.Add(dto)
}
