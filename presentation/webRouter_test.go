package presentation_test

import (
	"go-micro/container"
	"go-micro/tools"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIntegration(t *testing.T) {
	cont := container.NewContainer()
	tools.GetLogger().SetOutput(io.Discard)

	cont.DataBase.InitializeSqlite().Migrate().CreateSampleData()

	urls := []string{"/articles", "/articles/1", "/articleTypes", "/invoices", "/invoices/1"}

	for _, url := range urls {
		w := httptest.NewRecorder()
		cont.WebServer.Router.GetRouter().ServeHTTP(w, httptest.NewRequest("GET", url, nil))

		if w.Code != http.StatusOK {
			t.Error("Did not get expected HTTP status code for '"+url+"', got", w.Code)
			t.Fail()
		}

	}

}
