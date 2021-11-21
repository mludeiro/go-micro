package presentation

import (
	"go-micro/database"
	"go-micro/repository"
	"go-micro/service"
	"net/http"

	"github.com/gorilla/mux"
)

func getRouter() *mux.Router {
	sm := mux.NewRouter()

	// this is not necessary
	cont := ArticlesController{articleService: service.Article{Repository: &repository.Article{DB: database.GetDB()}}}

	sm.Methods(http.MethodGet).Path("/articles/{id:[0-9]+}").HandlerFunc(cont.GetArticle)
	sm.Methods(http.MethodGet).Path("/articles").HandlerFunc(cont.GetArticles)

	sm.Methods(http.MethodDelete).Path("/articles/{id:[0-9]+}").HandlerFunc(cont.DeleteArticle)
	sm.Methods(http.MethodPost).Path("/articles").HandlerFunc(cont.PostArticle)

	// Sometimes we only need handlers ;-)
	sm.Methods(http.MethodGet).Path("/articleTypes").HandlerFunc(GetArticleTypesHandler)

	invo := InvoiceController{InvoiceService: service.Invoice{Repository: &repository.Invoice{DB: database.GetDB()}}}
	sm.Methods(http.MethodGet).Path("/invoices/{id:[0-9]+}").HandlerFunc(invo.Get)
	sm.Methods(http.MethodGet).Path("/invoices").HandlerFunc(invo.GetInvoices)

	sm.Methods(http.MethodPost).Path("/invoices").HandlerFunc(invo.PostInvoice)

	// Just to log the calls, response code and time spent
	sm.Use(LogMiddleWare)

	return sm
}
