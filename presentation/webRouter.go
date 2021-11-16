package presentation

import (
	"net/http"

	"github.com/gorilla/mux"
)

func getRouter() *mux.Router {
	sm := mux.NewRouter()

	// this is not necessary
	cont := ArticlesController{}

	sm.Methods(http.MethodGet).Subrouter().Path("/articles/{id:[0-9]+}").Queries("expand", "{expand:[a-zA-Z0-9]+}").HandlerFunc(cont.GetArticle)
	//	sm.Methods(http.MethodGet).Subrouter().Path("/articles/{id:[0-9]+}").HandlerFunc(cont.GetArticle)
	sm.Methods(http.MethodGet).Subrouter().Path("/articles").Queries("expand", "{expand}").HandlerFunc(cont.GetArticles)

	sm.Methods(http.MethodDelete).Subrouter().HandleFunc("/articles/{id:[0-9]+}", cont.DeleteArticle)
	sm.Methods(http.MethodPost).Subrouter().HandleFunc("/articles", cont.PostArticle)

	// Sometimes we only need handlers ;-)
	sm.Methods(http.MethodGet).Subrouter().HandleFunc("/articleTypes", GetArticleTypesHandler)

	sm.Use(LogMiddleWare)

	return sm
}
