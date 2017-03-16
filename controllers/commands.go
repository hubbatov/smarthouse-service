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

//GetCommands returns json with all control commands in house
func GetCommands(w http.ResponseWriter, req *http.Request) {
	accessToken := req.Header.Get("Authorization")

	fl, userID := CheckAuthorization(accessToken)

	if fl && GetIntVar("user_id", req) == userID {
		w.WriteHeader(http.StatusOK)

		houseID := GetIntVar("house_id", req)

		json.NewEncoder(w).Encode(DBManager.commands(houseID))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "%s", "Please, login or register")
	}
}

//AddCommand creates new control command in house
func AddCommand(w http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")

	if !strings.Contains(contentType, "application/json") {
		errors.HandleError(errors.GenerateCustomError("Content-Type is not application/json"))
	}

	body, err := ioutil.ReadAll(req.Body)
	errors.HandleError(errors.ConvertCustomError(err))

	var controlCommand restapi.RESTCommand
	err = json.Unmarshal(body, &controlCommand)
	errors.HandleError(errors.ConvertCustomError(err))

	accessToken := req.Header.Get("Authorization")
	fl, userID := CheckAuthorization(accessToken)

	if fl && GetIntVar("user_id", req) == userID {

		houseID := GetIntVar("house_id", req)

		controlCommand.HouseID = houseID
		eArray := DBManager.addCommand(controlCommand)

		if len(eArray) > 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "%s", eArray)
		} else {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "%s", "Created command")
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "%s", "Please, login or register")
	}
}

//EditCommand edit existing control command in house
func EditCommand(w http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		errors.HandleError(errors.GenerateCustomError("Content-Type is not application/json"))
	}

	accessToken := req.Header.Get("Authorization")
	fl, userID := CheckAuthorization(accessToken)

	if fl && GetIntVar("user_id", req) == userID {

		body, err := ioutil.ReadAll(req.Body)
		errors.HandleError(errors.ConvertCustomError(err))

		var controlCommand restapi.RESTCommand
		err = json.Unmarshal(body, &controlCommand)
		errors.HandleError(errors.ConvertCustomError(err))

		commandID := GetIntVar("command_id", req)

		eArray := DBManager.editCommand(commandID, controlCommand)

		if len(eArray) > 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "%s", eArray)
		} else {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "%s %d", "Edited command", commandID)
		}

	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "%s", "Please, login or register")
	}
}

//RemoveCommand removes command from house
func RemoveCommand(w http.ResponseWriter, req *http.Request) {
	accessToken := req.Header.Get("Authorization")
	fl, userID := CheckAuthorization(accessToken)

	if fl && GetIntVar("user_id", req) == userID {

		commandID := GetIntVar("command_id", req)

		eArray := DBManager.removeCommand(commandID)

		if len(eArray) > 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "%s", eArray)
		} else {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "%s %d", "Removed command", commandID)
		}

	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "%s", "Please, login or register")
	}
}
