package controllers

import (
	"net/http"
	"strconv"

	"anybodyhere/errors"

	"github.com/gorilla/mux"
)

//GetIntVar returns int variable by id
func GetIntVar(id string, req *http.Request) int {
	vars := mux.Vars(req)
	result, err := strconv.Atoi(vars[id])
	errors.HandleError(errors.ConvertCustomError(err))
	return result
}
