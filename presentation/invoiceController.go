package presentation

import (
	"encoding/json"
	"go-micro/entity"
	"go-micro/service"
	"go-micro/tools"
	"net/http"
)

type InvoiceController struct {
	InvoiceService service.IInvoice
}

func (this InvoiceController) Get(rw http.ResponseWriter, r *http.Request) {
	id, _ := GetUIntParam(r, "id")
	article := this.InvoiceService.Get(id, GetExpand(r))

	if article == nil {
		rw.WriteHeader(http.StatusNotFound)
	} else {
		str, err := json.Marshal(article)

		if err == nil {
			rw.WriteHeader(http.StatusOK)
			rw.Write(str)
		} else {
			rw.WriteHeader(http.StatusInternalServerError)
			tools.GetLogger().Println(err)
		}
	}

}

func (this InvoiceController) GetInvoices(rw http.ResponseWriter, r *http.Request) {
	str, err := json.Marshal(this.InvoiceService.GetAll(GetExpand(r)))

	if err == nil {
		rw.WriteHeader(http.StatusOK)
		rw.Write(str)
	} else {
		rw.WriteHeader(http.StatusInternalServerError)
		tools.GetLogger().Println(err)
	}
}

func (this InvoiceController) PostInvoice(rw http.ResponseWriter, r *http.Request) {
	dto := &entity.Invoice{}
	err := json.NewDecoder(r.Body).Decode(dto)

	if err != nil {
		tools.GetLogger().Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(err.Error()))
		return
	}

	dto, err = this.InvoiceService.Add(dto)

	if err != nil {
		tools.GetLogger().Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
		return
	}

	str, err := json.Marshal(*dto)

	if err == nil {
		rw.WriteHeader(http.StatusCreated)
		rw.Write(str)
	} else {
		rw.WriteHeader(http.StatusInternalServerError)
		tools.GetLogger().Println(err)
	}
}
