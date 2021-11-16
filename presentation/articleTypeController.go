package presentation

import (
	"encoding/json"
	"go-micro/repository"
	"go-micro/tools"
	"net/http"
)

func GetArticleTypesHandler(rw http.ResponseWriter, r *http.Request) {
	str, err := json.Marshal(repository.GetArticleTypes())

	if err == nil {
		rw.WriteHeader(http.StatusOK)
		rw.Write(str)
	} else {
		rw.WriteHeader(http.StatusInternalServerError)
		tools.GetLogger().Println(err)
	}
}
