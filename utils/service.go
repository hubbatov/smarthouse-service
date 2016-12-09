package utils

import (
	"log"
	"net/http"

	"anybodyhere/controllers"

	"github.com/gorilla/mux"
)

// Service for data access
type Service struct {
}

//Run service
func (s *Service) Run() {
	log.Print("Starting service ...")

	controllers.CreateDb()

	router := mux.NewRouter()

	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/users", controllers.RegisterUser).Methods("POST")

	router.HandleFunc("/login", controllers.Login).Methods("POST")

	router.HandleFunc("/users/{user_id}/houses", controllers.GetHouses).Methods("GET")
	router.HandleFunc("/users/{user_id}/houses", controllers.AddHouse).Methods("POST")
	router.HandleFunc("/users/{user_id}/houses/{house_id}", controllers.EditHouse).Methods("PUT")
	router.HandleFunc("/users/{user_id}/houses/{house_id}", controllers.RemoveHouse).Methods("DELETE")

	router.HandleFunc("/users/{user_id}/houses/{house_id}/sensors", controllers.GetSensors).Methods("GET")
	router.HandleFunc("/users/{user_id}/houses/{house_id}/sensors", controllers.AddSensor).Methods("POST")
	router.HandleFunc("/users/{user_id}/houses/{house_id}/sensors/{sensor_id}", controllers.EditSensor).Methods("PUT")
	router.HandleFunc("/users/{user_id}/houses/{house_id}/sensors/{sensor_id}", controllers.RemoveSensor).Methods("DELETE")

	router.HandleFunc("/users/{user_id}/houses/{house_id}/sensors/{sensor_id}/sensordata", controllers.GetSensorData).Methods("GET")
	router.HandleFunc("/users/{user_id}/houses/{house_id}/sensors/{sensor_id}/sensordata", controllers.AddSensorData).Methods("POST")

	http.ListenAndServe(":12345", router)
}
