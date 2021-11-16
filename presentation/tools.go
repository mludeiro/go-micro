package presentation

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func GetUIntParam(r *http.Request, param string) (uint, error) {
	val, ok := mux.Vars(r)[param]

	if !ok {
		return 0, errors.New("Parameter not found")
	}

	// convert the value into an integer and return
	ival, err := strconv.Atoi(val)
	if err != nil {
		return 0, errors.New("Error parsing param")
	}

	return uint(ival), nil
}

func GetStringParam(r *http.Request, param string) (string, error) {
	vars := mux.Vars(r)

	val, exist := vars[param]

	if !exist {
		return "", errors.New("Parameter not found")
	}

	return val, nil
}

func GetArrayStringParam(r *http.Request, param string) ([]string, error) {
	val, err := GetStringParam(r, param)

	if err != nil {
		return []string{}, err
	}

	if len(strings.TrimSpace(val)) == 0 {
		return []string{}, nil
	}

	return strings.Split(strings.Replace(val, " ", "", -1), ","), nil
}

func GetExpand(r *http.Request) []string {
	val, _ := GetArrayStringParam(r, "expand")
	return val
}
