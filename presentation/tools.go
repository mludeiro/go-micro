package presentation

import (
	"errors"
	"go-micro/entity"
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

func GetStringQueryParam(r *http.Request, param string) (string, error) {
	val := r.URL.Query().Get(param)

	if len(val) == 0 {
		return "", errors.New("Parameter not found")
	}

	return val, nil
}

func GetArrayStringQueryParam(r *http.Request, param string) ([]string, error) {
	val, err := GetStringQueryParam(r, param)

	if err != nil {
		return []string{}, err
	}

	if len(strings.TrimSpace(val)) == 0 {
		return []string{}, nil
	}

	//return strings.Split(strings.Replace(val, " ", "", -1), ","), nil
	return strings.Split(val, ","), nil
}

func GetExpand(r *http.Request) []string {
	val, _ := GetArrayStringQueryParam(r, "expand")
	return val
}

func GetOrder(r *http.Request) []string {
	val, _ := GetArrayStringQueryParam(r, "order")
	return val
}

func GetCondition(r *http.Request) []entity.Condition {
	val, _ := GetArrayStringQueryParam(r, "condition")

	conditions := []entity.Condition{}

	for _, cond := range val {
		elems := strings.Split(cond, " ")
		if len(elems) != 3 {
			return []entity.Condition{}
		}
		conditions = append(conditions, entity.Condition{Field: elems[0], Comparator: elems[1], Value: elems[2]})
	}

	return conditions
}

func GetQuery(r *http.Request) entity.Query {
	return entity.Query{
		Fetchs:     GetExpand(r),
		Conditions: GetCondition(r),
		OrderBy:    GetOrder(r),
	}
}
