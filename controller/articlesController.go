package controller

import (
	"encoding/json"
	"go-micro/repository"
	"log"
	"net/http"
)

type ArticlesController struct {
	Logger *log.Logger
}

func (a *ArticlesController) GetArticlesHandler(rw http.ResponseWriter, r *http.Request) {
	str, err := json.Marshal(repository.GetArticles())

	if err == nil {
		rw.WriteHeader(http.StatusOK)
		rw.Write(str)
	} else {
		rw.WriteHeader(http.StatusInternalServerError)
		a.Logger.Println(err)
	}
}
