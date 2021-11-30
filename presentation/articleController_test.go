package presentation_test

import (
	"errors"
	"go-micro/entity"
	"go-micro/presentation"
	"net/http"
	"net/http/httptest"
	"testing"
)

type ServiceActicleMock struct {
	Error        error
	Article      *entity.Article
	ArticleArray []entity.Article
}

func (this ServiceActicleMock) Get(uint, []string) *entity.Article {
	return this.Article
}

func (this ServiceActicleMock) GetAll([]string) []entity.Article {
	return this.ArticleArray
}

func (this ServiceActicleMock) Add(*entity.Article) (*entity.Article, error) {
	return this.Article, this.Error
}

func (this ServiceActicleMock) Delete(uint) *entity.Article {
	return this.Article
}

func TestGet(t *testing.T) {
	w := httptest.NewRecorder()
	c := presentation.ArticleController{Service: ServiceActicleMock{Article: &entity.Article{}}}

	c.GetArticle(w, httptest.NewRequest("GET", "/", nil))

	if w.Code != http.StatusOK {
		t.Error("Did not get expected HTTP status code, got", w.Code)
		t.Fail()
	}
}

func TestGetFail(t *testing.T) {
	w := httptest.NewRecorder()
	c := presentation.ArticleController{Service: ServiceActicleMock{Error: errors.New("")}}

	c.GetArticle(w, httptest.NewRequest("GET", "/", nil))

	if w.Code != http.StatusNotFound {
		t.Error("Did not get expected HTTP status code, got", w.Code)
		t.Fail()
	}
}
