package presentation

import (
	"encoding/json"
	"go-micro/service"
	"go-micro/tools"
	"net/http"
)

type ArticleTypeController struct {
	Service service.IArticleType
}

func (this *ArticleTypeController) GetAll(rw http.ResponseWriter, r *http.Request) {
	data, err := this.Service.GetAll(GetExpand(r))
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
	} else {
		str, err := json.Marshal(data)
		if err == nil {
			rw.WriteHeader(http.StatusOK)
			rw.Write(str)
		} else {
			rw.WriteHeader(http.StatusInternalServerError)
			tools.GetLogger().Println(err)
		}
	}

}
