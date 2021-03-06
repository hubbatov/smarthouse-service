package utils

import (
	"log"
	"net/http"

	"smarthouse-service/controllers"

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

	router.HandleFunc("/user", controllers.GetUser).Methods("GET")
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

	router.HandleFunc("/users/{user_id}/houses/{house_id}/commands", controllers.GetCommands).Methods("GET")
	router.HandleFunc("/users/{user_id}/houses/{house_id}/commands", controllers.AddCommand).Methods("POST")
	router.HandleFunc("/users/{user_id}/houses/{house_id}/commands/{command_id}", controllers.EditCommand).Methods("PUT")
	router.HandleFunc("/users/{user_id}/houses/{house_id}/commands/{command_id}", controllers.RemoveCommand).Methods("DELETE")
	router.HandleFunc("/users/{user_id}/houses/{house_id}/commands/{command_id}", controllers.RemoveCommand).Methods("DELETE")
	router.HandleFunc("/commands/do", controllers.RunCommand).Methods("POST")

	router.HandleFunc("/users/{user_id}/houses/{house_id}/sensors/{sensor_id}/sensordata", controllers.GetSensorData).Methods("PUT")
	router.HandleFunc("/users/{user_id}/houses/{house_id}/sensors/{sensor_id}/sensordata", controllers.AddSensorData).Methods("POST")
	router.HandleFunc("/sensordata/{sensor_tag}", controllers.AddSensorDataByTag).Methods("POST")

	http.ListenAndServe(":12345", router)
}
