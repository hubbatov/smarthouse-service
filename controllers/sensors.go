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

//GetSensors returns json with all sensors in house
func GetSensors(w http.ResponseWriter, req *http.Request) {
	accessToken := req.Header.Get("Authorization")

	fl, userID := CheckAuthorization(accessToken)

	if fl && GetIntVar("user_id", req) == userID {
		w.WriteHeader(http.StatusOK)

		houseID := GetIntVar("house_id", req)

		json.NewEncoder(w).Encode(DBManager.sensors(houseID))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "%s", "Please, login or register")
	}
}

//AddSensor creates new sensor in house
func AddSensor(w http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")

	if !strings.Contains(contentType, "application/json") {
		errors.HandleError(errors.GenerateCustomError("Content-Type is not application/json"))
	}

	body, err := ioutil.ReadAll(req.Body)
	errors.HandleError(errors.ConvertCustomError(err))

	var sensordata restapi.RESTSensor
	err = json.Unmarshal(body, &sensordata)
	errors.HandleError(errors.ConvertCustomError(err))

	accessToken := req.Header.Get("Authorization")
	fl, userID := CheckAuthorization(accessToken)

	if fl && GetIntVar("user_id", req) == userID {

		houseID := GetIntVar("house_id", req)

		sensordata.HouseID = houseID
		eArray := DBManager.addSensor(sensordata)

		if len(eArray) > 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "%s", eArray)
		} else {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "%s", "Created sensor")
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "%s", "Please, login or register")
	}
}

//EditSensor edit existing sensor in house
func EditSensor(w http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		errors.HandleError(errors.GenerateCustomError("Content-Type is not application/json"))
	}

	accessToken := req.Header.Get("Authorization")
	fl, userID := CheckAuthorization(accessToken)

	if fl && GetIntVar("user_id", req) == userID {

		body, err := ioutil.ReadAll(req.Body)
		errors.HandleError(errors.ConvertCustomError(err))

		var sensordata restapi.RESTSensor
		err = json.Unmarshal(body, &sensordata)
		errors.HandleError(errors.ConvertCustomError(err))

		sensorID := GetIntVar("sensor_id", req)

		eArray := DBManager.editSensor(sensorID, sensordata)

		if len(eArray) > 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "%s", eArray)
		} else {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "%s %d", "Edited sensor", sensorID)
		}

	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "%s", "Please, login or register")
	}
}

//RemoveSensor removes sensor from house
func RemoveSensor(w http.ResponseWriter, req *http.Request) {
	accessToken := req.Header.Get("Authorization")
	fl, userID := CheckAuthorization(accessToken)

	if fl && GetIntVar("user_id", req) == userID {

		sensorID := GetIntVar("sensor_id", req)

		eArray := DBManager.removeSensor(sensorID)

		if len(eArray) > 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "%s", eArray)
		} else {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "%s %d", "Removed sensor", sensorID)
		}

	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "%s", "Please, login or register")
	}
}

//GetSensorData returns json with all sensor data
func GetSensorData(w http.ResponseWriter, req *http.Request) {
	accessToken := req.Header.Get("Authorization")

	fl, userID := CheckAuthorization(accessToken)

	if fl && GetIntVar("user_id", req) == userID {
		w.WriteHeader(http.StatusOK)

		sensorID := GetIntVar("sensor_id", req)

		json.NewEncoder(w).Encode(DBManager.sensordata(sensorID))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "%s", "Please, login or register")
	}
}

//AddSensorData adds sensor data
func AddSensorData(w http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		errors.HandleError(errors.GenerateCustomError("Content-Type is not application/json"))
	}

	accessToken := req.Header.Get("Authorization")
	fl, userID := CheckAuthorization(accessToken)

	if fl && GetIntVar("user_id", req) == userID {

		body, err := ioutil.ReadAll(req.Body)
		errors.HandleError(errors.ConvertCustomError(err))

		var sensordata restapi.RESTSensorData
		err = json.Unmarshal(body, &sensordata)
		errors.HandleError(errors.ConvertCustomError(err))

		sensorID := GetIntVar("sensor_id", req)
		sensordata.SensorID = sensorID

		eArray := DBManager.addSensorData(sensordata)

		if len(eArray) > 0 {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "%s", eArray)
		} else {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "%s %d", "Added sensor data to sensor", sensorID)
		}

	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "%s", "Please, login or register")
	}
}

//AddSensorDataByTag adds sensor data using sensor tag
func AddSensorDataByTag(w http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		errors.HandleError(errors.GenerateCustomError("Content-Type is not application/json"))
	}

	accessToken := req.Header.Get("Authorization")
	fl, userID := CheckAuthorization(accessToken)

	if fl {

		body, err := ioutil.ReadAll(req.Body)
		errors.HandleError(errors.ConvertCustomError(err))

		var sensordata restapi.RESTSensorData
		err = json.Unmarshal(body, &sensordata)
		errors.HandleError(errors.ConvertCustomError(err))

		sensorTag := GetStringVar("sensor_tag", req)

		sensorID := DBManager.sensorID(sensorTag)
		sensordata.SensorID = sensorID

		if DBManager.userID(sensorID) == userID {
			eArray := DBManager.addSensorData(sensordata)

			if len(eArray) > 0 {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "%s", eArray)
			} else {
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, "%s %d", "Added sensor data to sensor", sensorID)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "%s", "This sensor is not yours")
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "%s", "Please, login or register")
	}
}
