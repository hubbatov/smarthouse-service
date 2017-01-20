package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"smarthouse-service/errors"
	"smarthouse-service/restapi"

	"strings"

	"fmt"
)

//GetUsers returns json with all users in system
func GetUsers(w http.ResponseWriter, req *http.Request) {
	accessToken := req.Header.Get("Authorization")

	fl, _ := CheckAuthorization(accessToken)

	if fl {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(DBManager.users())
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "%s", "Please, login or register")
	}
}

//GetUser returns json with authorized user data
func GetUser(w http.ResponseWriter, req *http.Request) {
	accessToken := req.Header.Get("Authorization")

	fl, userID := CheckAuthorization(accessToken)

	if fl {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(GetUserByID(userID))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "%s", "Please, login or register")
	}
}

//RegisterUser creates new user in system
func RegisterUser(w http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")

	if !strings.Contains(contentType, "application/json") {
		errors.HandleError(errors.GenerateCustomError("Content-Type is not application/json"))
	}

	body, err := ioutil.ReadAll(req.Body)
	errors.HandleError(errors.ConvertCustomError(err))

	var userdata restapi.RESTUser
	err = json.Unmarshal(body, &userdata)
	errors.HandleError(errors.ConvertCustomError(err))

	eArray := DBManager.createUser(userdata)

	if len(eArray) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", eArray)
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", "Registered")
	}
}
