package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"anybodyhere/errors"
	"anybodyhere/models"
	"anybodyhere/restapi"

	"strings"

	"fmt"
)

// UsersController provides methods for users
type UsersController struct {
}

//GetUsers returns json with all users in system
func (c *UsersController) GetUsers(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(DBManager.users())
}

//RegisterUser creates new user in system
func (c *UsersController) RegisterUser(w http.ResponseWriter, req *http.Request) {
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

//Login to service and get unique cached token
func (c *UsersController) Login(w http.ResponseWriter, req *http.Request) {

	contentType := req.Header.Get("Content-Type")

	if !strings.Contains(contentType, "application/json") {
		errors.HandleError(errors.GenerateCustomError("Content-Type is not application/json"))
	}

	body, err := ioutil.ReadAll(req.Body)
	errors.HandleError(errors.ConvertCustomError(err))

	var userdata restapi.RESTUser
	err = json.Unmarshal(body, &userdata)
	errors.HandleError(errors.ConvertCustomError(err))

	user := getUser(userdata.Login)

	if user.Login == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "%s", "User not found")
	} else {

		updatePasswordHash(&user)

		err = bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(userdata.Password))
		errors.HandleError(errors.ConvertCustomError(err))

		if err == nil {
			w.WriteHeader(http.StatusAccepted)
			fmt.Fprintf(w, "%s", user.HashedPassword)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "%s", "Incorrect password")
		}
	}
}

func updatePasswordHash(user *models.User) {
	if user != nil {
		user.HashedPassword, _ = bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	}
}

func getUser(login string) (user models.User) {
	DBManager.dataBase.Where("login = ?", login).First(&user)
	return
}
