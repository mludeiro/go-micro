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

// exposing all the repository methos with no filter
type Invoice struct {
	repository.IInvoice
}
