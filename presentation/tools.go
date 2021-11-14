package presentation

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetIntParam(r *http.Request, param string) int {
	vars := mux.Vars(r)

	// convert the value into an integer and return
	val, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		panic(err)
	}

	return val
}
