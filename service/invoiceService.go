package service

import (
	"go-micro/entity"
	"go-micro/repository"
)

type IInvoiceService interface {
	Get(uint, []string) (*entity.Invoice, error)
	GetAll([]string) ([]entity.Invoice, error)
	Add(*entity.Invoice) (*entity.Invoice, error)
}

// exposing all the repository methos with no filter
type Invoice struct {
	repository.IInvoiceRepository
}
