package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"smarthouse-service/errors"
	"smarthouse-service/models"
	"smarthouse-service/restapi"

	"strings"

	"encoding/base64"
	"fmt"
)

//Login to service and get unique cached token
func Login(w http.ResponseWriter, req *http.Request) {

	contentType := req.Header.Get("Content-Type")

	if !strings.Contains(contentType, "application/json") {
		errors.HandleError(errors.GenerateCustomError("Content-Type is not application/json"))
	}

	body, err := ioutil.ReadAll(req.Body)
	errors.HandleError(errors.ConvertCustomError(err))

	var userdata restapi.RESTUser
	err = json.Unmarshal(body, &userdata)
	errors.HandleError(errors.ConvertCustomError(err))

	fl, _ := CheckLoginPassword(userdata.Login, userdata.Password)

	if fl {
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
func CheckAuthorization(authdata string) (bool, int) {
	data := authdata
	data = strings.Replace(data, "Basic ", "", -1)
	login, password := getUserData(data)
	return CheckLoginPassword(login, password)
}

//CheckLoginPassword validates pare login+password
func CheckLoginPassword(login, password string) (bool, int) {
	user := getUser(login)

	if user.Login == "" {
		return false, 0
	}
	updatePasswordHash(&user)

	err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password))
	errors.HandleError(errors.ConvertCustomError(err))

	if err == nil {
		return true, user.ID
	}
	return false, 0
}
