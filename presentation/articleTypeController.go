package presentation

import (
	"encoding/json"
	"go-micro/service"
	"go-micro/tools"
	"net/http"
)

type ArticleTypeController struct {
	Service *service.ArticleType
}

func (this *ArticleTypeController) GetAll(rw http.ResponseWriter, r *http.Request) {
	str, err := json.Marshal(this.Service.GetArticleTypes(GetExpand(r)))

	if err == nil {
		rw.WriteHeader(http.StatusOK)
		rw.Write(str)
	} else {
		rw.WriteHeader(http.StatusInternalServerError)
		tools.GetLogger().Println(err)
	}
}
