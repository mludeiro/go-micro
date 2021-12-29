package presentation

import (
	"net/http"

	"github.com/gorilla/mux"
)

type WebRouter struct {
	muxRuter              *mux.Router
	ArticleController     ArticleController
	ArticleTypeController ArticleTypeController
	InvoiceController     InvoiceController
}

func (router *WebRouter) GetRouter() *mux.Router {
	if router.muxRuter != nil {
		return router.muxRuter
	}

	sm := mux.NewRouter()

	sm.Methods(http.MethodGet).Path("/articles/{id:[0-9]+}").HandlerFunc(router.ArticleController.GetArticle)
	sm.Methods(http.MethodGet).Path("/articles").HandlerFunc(router.ArticleController.GetArticles)

	sm.Methods(http.MethodDelete).Path("/articles/{id:[0-9]+}").HandlerFunc(router.ArticleController.DeleteArticle)
	sm.Methods(http.MethodPost).Path("/articles").HandlerFunc(router.ArticleController.PostArticle)

	// Sometimes we only need handlers ;-)
	sm.Methods(http.MethodGet).Path("/articleTypes").HandlerFunc(router.ArticleTypeController.GetAll)

	//invo := InvoiceController{InvoiceService: service.Invoice{Repository: &repository.Invoice{DB: database.GetDB()}}}
	sm.Methods(http.MethodGet).Path("/invoices/{id:[0-9]+}").HandlerFunc(router.InvoiceController.Get)
	sm.Methods(http.MethodGet).Path("/invoices").HandlerFunc(router.InvoiceController.GetInvoices)
	sm.Methods(http.MethodPost).Path("/invoices").HandlerFunc(router.InvoiceController.PostInvoice)

	// Just to log the calls, response code and time spent
	sm.Use(LogMiddleWare)

	router.muxRuter = sm
	return router.muxRuter
}
