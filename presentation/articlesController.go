package presentation

import (
	"encoding/json"
	"go-micro/repository"
	"go-micro/tools"
	"net/http"
)

type ArticlesController struct {
}

func (a *ArticlesController) GetArticle(rw http.ResponseWriter, r *http.Request) {
	article := repository.GetArticle(uint(GetIntParam(r, "id")))

	if article == nil {
		rw.WriteHeader(http.StatusNotFound)
	} else {
		str, err := json.Marshal(*article)

		if err == nil {
			rw.WriteHeader(http.StatusOK)
			rw.Write(str)
		} else {
			rw.WriteHeader(http.StatusInternalServerError)
			tools.GetLogger().Println(err)
		}
	}

}

func (a *ArticlesController) DeleteArticle(rw http.ResponseWriter, r *http.Request) {
	article := repository.DeleteArticle(uint(GetIntParam(r, "id")))

	if article == nil {
		rw.WriteHeader(http.StatusNotFound)
	} else {
		str, err := json.Marshal(*article)

		if err == nil {
			rw.WriteHeader(http.StatusOK)
			rw.Write(str)
		} else {
			rw.WriteHeader(http.StatusInternalServerError)
			tools.GetLogger().Println(err)
		}
	}

}

func (a *ArticlesController) GetArticles(rw http.ResponseWriter, r *http.Request) {
	str, err := json.Marshal(repository.GetArticles())

	if err == nil {
		rw.WriteHeader(http.StatusOK)
		rw.Write(str)
	} else {
		rw.WriteHeader(http.StatusInternalServerError)
		tools.GetLogger().Println(err)
	}
}

func (a *ArticlesController) PostArticle(rw http.ResponseWriter, r *http.Request) {
	dto := &repository.Article{}
	err := json.NewDecoder(r.Body).Decode(dto)

	if err != nil {
		tools.GetLogger().Println(err)
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
		tools.GetLogger().Println(err)
	}
}
