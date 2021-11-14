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

func (a *ArticlesController) GetArticles(rw http.ResponseWriter, r *http.Request) {
	str, err := json.Marshal(repository.GetArticles())

	if err == nil {
		rw.WriteHeader(http.StatusOK)
		rw.Write(str)
	} else {
		rw.WriteHeader(http.StatusInternalServerError)
		a.Logger.Println(err)
	}
}

func (a *ArticlesController) PostArticle(rw http.ResponseWriter, r *http.Request) {
	dto := &repository.Article{}
	err := json.NewDecoder(r.Body).Decode(dto)

	if err != nil {
		a.Logger.Println(err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	repository.AddArticle(dto)

	str, err := json.Marshal(*dto)

	if err == nil {
		rw.WriteHeader(http.StatusCreated)
		rw.Write(str)
	} else {
		rw.WriteHeader(http.StatusInternalServerError)
		a.Logger.Println(err)
	}
}
