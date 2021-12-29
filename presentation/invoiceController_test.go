package presentation_test

import (
	"bytes"
	"errors"
	"go-micro/entity"
	"go-micro/presentation"
	"net/http"
	"net/http/httptest"
	"testing"
)

type ServiceInvoiceMock struct {
	Error        error
	Invoice      *entity.Invoice
	InvoiceArray []entity.Invoice
}

func (mock *ServiceInvoiceMock) Get(uint, []string) (*entity.Invoice, error) {
	return mock.Invoice, mock.Error
}

func (mock *ServiceInvoiceMock) GetAll([]string) ([]entity.Invoice, error) {
	return mock.InvoiceArray, mock.Error
}

func (mock *ServiceInvoiceMock) Add(*entity.Invoice) (*entity.Invoice, error) {
	return mock.Invoice, mock.Error
}

func TestGetAllInvoices(t *testing.T) {
	w := httptest.NewRecorder()
	c := presentation.InvoiceController{InvoiceService: &ServiceInvoiceMock{Invoice: &entity.Invoice{}}}

	c.GetInvoices(w, httptest.NewRequest("GET", "/", nil))

	if w.Code != http.StatusOK {
		t.Error("Did not get expected HTTP status code, got", w.Code)
		t.Fail()
	}
}

func TestGetAllInvoiceNotFound(t *testing.T) {
	w := httptest.NewRecorder()
	c := presentation.InvoiceController{InvoiceService: &ServiceInvoiceMock{Invoice: &entity.Invoice{}}}

	c.GetInvoices(w, httptest.NewRequest("GET", "/", nil))

	if w.Code != http.StatusOK {
		t.Error("Did not get expected HTTP status code, got", w.Code)
		t.Fail()
	}
}
func TestGetAllInvoicesFail(t *testing.T) {
	w := httptest.NewRecorder()
	c := presentation.InvoiceController{InvoiceService: &ServiceInvoiceMock{Error: errors.New("")}}

	c.GetInvoices(w, httptest.NewRequest("GET", "/", nil))

	if w.Code != http.StatusInternalServerError {
		t.Error("Did not get expected HTTP status code, got", w.Code)
		t.Fail()
	}
}

func TestPostInvoiceBadRequest(t *testing.T) {
	w := httptest.NewRecorder()
	c := presentation.InvoiceController{InvoiceService: &ServiceInvoiceMock{Error: errors.New("")}}

	c.PostInvoice(w, httptest.NewRequest("POST", "/", nil))

	if w.Code != http.StatusBadRequest {
		t.Error("Did not get expected HTTP status code, got", w.Code)
		t.Fail()
	}

	c.PostInvoice(w, httptest.NewRequest("POST", "/", bytes.NewBuffer([]byte("{}"))))
}

func TestPostInvoiceServerError(t *testing.T) {
	w := httptest.NewRecorder()
	c := presentation.InvoiceController{InvoiceService: &ServiceInvoiceMock{Error: errors.New("")}}

	c.PostInvoice(w, httptest.NewRequest("POST", "/", bytes.NewBuffer([]byte("{}"))))

	if w.Code != http.StatusInternalServerError {
		t.Error("Did not get expected HTTP status code, got", w.Code)
		t.Fail()
	}

}
