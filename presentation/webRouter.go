package presentation

import (
	"net/http"

	"github.com/gorilla/mux"
)

func getRouter() *mux.Router {
	sm := mux.NewRouter()

	cont := ArticlesController{}

	sm.Methods(http.MethodGet).Subrouter().HandleFunc("/articles/{id:[0-9]+}", cont.GetArticle)
	sm.Methods(http.MethodDelete).Subrouter().HandleFunc("/articles/{id:[0-9]+}", cont.DeleteArticle)
	sm.Methods(http.MethodGet).Subrouter().HandleFunc("/articles", cont.GetArticles)
	sm.Methods(http.MethodPost).Subrouter().HandleFunc("/articles", cont.PostArticle)

	sm.Use(LogMiddleWare)

	return sm
}
