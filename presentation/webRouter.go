package presentation

import (
	"net/http"

	"github.com/gorilla/mux"
)

func getRouter() *mux.Router {
	sm := mux.NewRouter()

	// this is not necessary
	cont := ArticlesController{}

	sm.Methods(http.MethodGet).Path("/articles/{id:[0-9]+}").HandlerFunc(cont.GetArticle)
	sm.Methods(http.MethodGet).Path("/articles").HandlerFunc(cont.GetArticles)

	sm.Methods(http.MethodDelete).Path("/articles/{id:[0-9]+}").HandlerFunc(cont.DeleteArticle)
	sm.Methods(http.MethodPost).Path("/articles").HandlerFunc(cont.PostArticle)

	// Sometimes we only need handlers ;-)
	sm.Methods(http.MethodGet).Path("/articleTypes").HandlerFunc(GetArticleTypesHandler)

	sm.Use(LogMiddleWare)

	return sm
}
