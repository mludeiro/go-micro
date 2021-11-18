package presentation

import (
	"encoding/json"
	"go-micro/entity"
	"go-micro/service"
	"go-micro/tools"
	"net/http"
)

type ArticlesController struct {
}

func (a *ArticlesController) GetArticle(rw http.ResponseWriter, r *http.Request) {
	id, _ := GetUIntParam(r, "id")
	article := service.GetArticle(id, GetExpand(r))

	if article == nil {
		rw.WriteHeader(http.StatusNotFound)
	} else {
		str, err := json.Marshal(article)

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
	article := service.DeleteArticle(id)

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
	str, err := json.Marshal(service.GetArticles(GetExpand(r)))

	if err == nil {
		rw.WriteHeader(http.StatusOK)
		rw.Write(str)
	} else {
		rw.WriteHeader(http.StatusInternalServerError)
		tools.GetLogger().Println(err)
	}
}

func (a *ArticlesController) PostArticle(rw http.ResponseWriter, r *http.Request) {
	dto := &entity.Article{}
	err := json.NewDecoder(r.Body).Decode(dto)

	if err != nil {
		tools.GetLogger().Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(err.Error()))
		return
	}

	dto, err = service.AddArticle(dto)

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
