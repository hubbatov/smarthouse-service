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

	"encoding/base64"
	"fmt"
)

// UsersController provides methods for users
type UsersController struct {
}

//GetUsers returns json with all users in system
func (c *UsersController) GetUsers(w http.ResponseWriter, req *http.Request) {
	accessToken := req.Header.Get("Authorization")

	if CheckAuthorization(accessToken) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(DBManager.users())
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "%s", "Please, login or register")
	}
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

	if CheckLoginPassword(userdata.Login, userdata.Password) {
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(w, "%s", "Authorized")
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "%s", "Login or password incorrect")
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

func getUserData(authdata string) (user, password string) {
	d, err := base64.StdEncoding.DecodeString(authdata)

	errors.HandleError(errors.ConvertCustomError(err))

	data := strings.Split(string(d), ":")

	if len(data) == 2 {
		return data[0], data[1]
	}

	return "", ""
}

//CheckAuthorization validates hash in Authorization header
func CheckAuthorization(authdata string) bool {
	login, password := getUserData(authdata)
	return CheckLoginPassword(login, password)
}

//checkLoginPassword validates pare login+password
func CheckLoginPassword(login, password string) bool {
	user := getUser(login)

	if user.Login == "" {
		return false
	}
	updatePasswordHash(&user)

	err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password))
	errors.HandleError(errors.ConvertCustomError(err))

	if err == nil {
		return true
	}
	return false
}
