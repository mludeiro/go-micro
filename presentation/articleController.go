package presentation

import (
	"encoding/json"
	"go-micro/entity"
	"go-micro/service"
	"go-micro/tools"
	"net/http"
)

type ArticleController struct {
	Service service.IActicleService
}

func (cont *ArticleController) GetArticle(rw http.ResponseWriter, r *http.Request) {
	id, _ := GetUIntParam(r, "id")
	article, err := cont.Service.Get(id, GetExpand(r))

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
	} else if article == nil {
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

func (cont *ArticleController) DeleteArticle(rw http.ResponseWriter, r *http.Request) {
	id, _ := GetUIntParam(r, "id")
	article, err := cont.Service.Delete(id)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
	} else if article == nil {
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

func (cont *ArticleController) GetArticles(rw http.ResponseWriter, r *http.Request) {
	query := GetRepositoryQuery(r)
	lst, err := cont.Service.GetAll(&query)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
	} else {
		str, err := json.Marshal(lst)

		if err == nil {
			rw.WriteHeader(http.StatusOK)
			rw.Write(str)
		} else {
			rw.WriteHeader(http.StatusInternalServerError)
			tools.GetLogger().Println(err)
		}
	}
}

func (cont *ArticleController) PostArticle(rw http.ResponseWriter, r *http.Request) {
	dto := &entity.Article{}
	err := json.NewDecoder(r.Body).Decode(dto)

	if err != nil {
		tools.GetLogger().Println(err)
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(err.Error()))
		return
	}

	dto, err = cont.Service.Add(dto)

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
