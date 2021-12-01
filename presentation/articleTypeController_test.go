package presentation_test

import (
	"errors"
	"go-micro/entity"
	"go-micro/presentation"
	"net/http"
	"net/http/httptest"
	"testing"
)

type ServiceActicleTypeMock struct {
	Error        error
	ArticleType  *entity.ArticleType
	ArticleArray []entity.ArticleType
}

func (this ServiceActicleTypeMock) Get(uint, []string) (*entity.ArticleType, error) {
	return this.ArticleType, this.Error
}

func (this ServiceActicleTypeMock) GetAll([]string) ([]entity.ArticleType, error) {
	return this.ArticleArray, this.Error
}

func TestGetAllArticleType(t *testing.T) {
	w := httptest.NewRecorder()
	c := presentation.ArticleTypeController{Service: ServiceActicleTypeMock{ArticleType: &entity.ArticleType{}}}

	c.GetAll(w, httptest.NewRequest("GET", "/", nil))

	if w.Code != http.StatusOK {
		t.Error("Did not get expected HTTP status code, got", w.Code)
		t.Fail()
	}
}

func TestGetAllArticleTypeNotFound(t *testing.T) {
	w := httptest.NewRecorder()
	c := presentation.ArticleTypeController{Service: ServiceActicleTypeMock{}}

	c.GetAll(w, httptest.NewRequest("GET", "/", nil))

	if w.Code != http.StatusOK {
		t.Error("Did not get expected HTTP status code, got", w.Code)
		t.Fail()
	}
}
func TestGetAllArticleTypeFail(t *testing.T) {
	w := httptest.NewRecorder()
	c := presentation.ArticleTypeController{Service: ServiceActicleTypeMock{Error: errors.New("")}}

	c.GetAll(w, httptest.NewRequest("GET", "/", nil))

	if w.Code != http.StatusInternalServerError {
		t.Error("Did not get expected HTTP status code, got", w.Code)
		t.Fail()
	}
}
