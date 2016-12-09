package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"anybodyhere/errors"
	"anybodyhere/restapi"

	"strings"

	"fmt"

	"strconv"

	"github.com/gorilla/mux"
)

//GetHouses returns json with all houses of user represented by Authorization
func GetHouses(w http.ResponseWriter, req *http.Request) {
	accessToken := req.Header.Get("Authorization")

	fl, userID := CheckAuthorization(accessToken)

	if fl {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(DBManager.houses(userID))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "%s", "Please, login or register")
	}
}

//AddHouse creates new house in system
func AddHouse(w http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")

	if !strings.Contains(contentType, "application/json") {
		errors.HandleError(errors.GenerateCustomError("Content-Type is not application/json"))
	}

	body, err := ioutil.ReadAll(req.Body)
	errors.HandleError(errors.ConvertCustomError(err))

	var housedata restapi.RESTHouse
	err = json.Unmarshal(body, &housedata)
	errors.HandleError(errors.ConvertCustomError(err))

	accessToken := req.Header.Get("Authorization")
	fl, userID := CheckAuthorization(accessToken)

	if fl {
		housedata.UserID = userID
		eArray := DBManager.addHouse(housedata)

		if len(eArray) > 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "%s", eArray)
		} else {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "%s", "Created house")
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "%s", "Please, login or register")
	}
}

//EditHouse edit existing house in system
func EditHouse(w http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		errors.HandleError(errors.GenerateCustomError("Content-Type is not application/json"))
	}

	accessToken := req.Header.Get("Authorization")
	fl, _ := CheckAuthorization(accessToken)

	if fl {

		body, err := ioutil.ReadAll(req.Body)
		errors.HandleError(errors.ConvertCustomError(err))

		var housedata restapi.RESTHouse
		err = json.Unmarshal(body, &housedata)
		errors.HandleError(errors.ConvertCustomError(err))

		vars := mux.Vars(req)
		houseID, err := strconv.Atoi(vars["id"])
		errors.HandleError(errors.ConvertCustomError(err))

		eArray := DBManager.editHouse(houseID, housedata)

		if len(eArray) > 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "%s", eArray)
		} else {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "%s %s", "Edited house", vars["id"])
		}

	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "%s", "Please, login or register")
	}
}

//RemoveHouse removes house from system
func RemoveHouse(w http.ResponseWriter, req *http.Request) {
	accessToken := req.Header.Get("Authorization")
	fl, _ := CheckAuthorization(accessToken)

	if fl {
		vars := mux.Vars(req)
		houseID, _ := strconv.Atoi(vars["id"])

		eArray := DBManager.removeHouse(houseID)

		if len(eArray) > 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "%s", eArray)
		} else {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "%s %s", "Removed house", vars["id"])
		}

	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "%s", "Please, login or register")
	}
}
