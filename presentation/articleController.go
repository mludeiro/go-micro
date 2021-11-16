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
	id, _ := GetUIntParam(r, "id")
	article := repository.GetArticle(id, GetExpand(r))

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
	id, _ := GetUIntParam(r, "id")
	article := repository.DeleteArticle(id)

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
	str, err := json.Marshal(repository.GetArticles(GetExpand(r)))

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
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(err.Error()))
		return
	}

	err = repository.AddArticle(dto)

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
