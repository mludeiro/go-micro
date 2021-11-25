package presentation

import (
	"net/http"

	"github.com/gorilla/mux"
)

type WebRouter struct {
	router                *mux.Router
	ArticleController     ArticleController
	ArticleTypeController ArticleTypeController
	InvoiceController     InvoiceController
}

func (this *WebRouter) getRouter() *mux.Router {
	sm := mux.NewRouter()

	// this is not necessary
	sm.Methods(http.MethodGet).Path("/articles/{id:[0-9]+}").HandlerFunc(this.ArticleController.GetArticle)
	sm.Methods(http.MethodGet).Path("/articles").HandlerFunc(this.ArticleController.GetArticles)

	sm.Methods(http.MethodDelete).Path("/articles/{id:[0-9]+}").HandlerFunc(this.ArticleController.DeleteArticle)
	sm.Methods(http.MethodPost).Path("/articles").HandlerFunc(this.ArticleController.PostArticle)

	// Sometimes we only need handlers ;-)
	sm.Methods(http.MethodGet).Path("/articleTypes").HandlerFunc(this.ArticleTypeController.GetAll)

	//invo := InvoiceController{InvoiceService: service.Invoice{Repository: &repository.Invoice{DB: database.GetDB()}}}
	sm.Methods(http.MethodGet).Path("/invoices/{id:[0-9]+}").HandlerFunc(this.InvoiceController.Get)
	sm.Methods(http.MethodGet).Path("/invoices").HandlerFunc(this.InvoiceController.GetInvoices)
	sm.Methods(http.MethodPost).Path("/invoices").HandlerFunc(this.InvoiceController.PostInvoice)

	// Just to log the calls, response code and time spent
	sm.Use(LogMiddleWare)

	return sm
}
