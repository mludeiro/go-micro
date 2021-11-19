package service

import (
	"go-micro/entity"
	"go-micro/repository"
)

var GetInvoice func(uint, []string) *entity.Invoice = repository.GetInvoice

var GetInvoices func([]string) []entity.Invoice = repository.GetInvoices

var AddInvoice func(*entity.Invoice) (*entity.Invoice, error) = repository.AddInvoice
