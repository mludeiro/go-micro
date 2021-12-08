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

func (this ServiceActicleMock) Get(uint, []string) (*entity.Article, error) {
	return this.Article, this.Error
}

func (this ServiceActicleMock) GetAll(query entity.Query) (entity.ArticleResultSet, error) {
	return entity.ArticleResultSet{Data: this.ArticleArray}, this.Error
}

func (this ServiceActicleMock) Add(*entity.Article) (*entity.Article, error) {
	return this.Article, this.Error
}

func (this ServiceActicleMock) Delete(uint) (*entity.Article, error) {
	return this.Article, this.Error
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

func TestGetNotFound(t *testing.T) {
	w := httptest.NewRecorder()
	c := presentation.ArticleController{Service: ServiceActicleMock{}}

	c.GetArticle(w, httptest.NewRequest("GET", "/", nil))

	if w.Code != http.StatusNotFound {
		t.Error("Did not get expected HTTP status code, got", w.Code)
		t.Fail()
	}
}
func TestGetFail(t *testing.T) {
	w := httptest.NewRecorder()
	c := presentation.ArticleController{Service: ServiceActicleMock{Error: errors.New("")}}

	c.GetArticle(w, httptest.NewRequest("GET", "/", nil))

	if w.Code != http.StatusInternalServerError {
		t.Error("Did not get expected HTTP status code, got", w.Code)
		t.Fail()
	}
}

func TestGetAll(t *testing.T) {
	w := httptest.NewRecorder()
	c := presentation.ArticleController{Service: ServiceActicleMock{Article: &entity.Article{}}}

	c.GetArticles(w, httptest.NewRequest("GET", "/", nil))

	if w.Code != http.StatusOK {
		t.Error("Did not get expected HTTP status code, got", w.Code)
		t.Fail()
	}
}

func TestGetAllNotFound(t *testing.T) {
	w := httptest.NewRecorder()
	c := presentation.ArticleController{Service: ServiceActicleMock{}}

	c.GetArticles(w, httptest.NewRequest("GET", "/", nil))

	if w.Code != http.StatusOK {
		t.Error("Did not get expected HTTP status code, got", w.Code)
		t.Fail()
	}
}
func TestGetAllFail(t *testing.T) {
	w := httptest.NewRecorder()
	c := presentation.ArticleController{Service: ServiceActicleMock{Error: errors.New("")}}

	c.GetArticles(w, httptest.NewRequest("GET", "/", nil))

	if w.Code != http.StatusInternalServerError {
		t.Error("Did not get expected HTTP status code, got", w.Code)
		t.Fail()
	}
}
