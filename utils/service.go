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

	router.HandleFunc("/houses", controllers.GetHouses).Methods("GET")
	router.HandleFunc("/houses", controllers.AddHouse).Methods("POST")
	router.HandleFunc("/houses/{id}", controllers.EditHouse).Methods("PUT")
	router.HandleFunc("/houses/{id}", controllers.RemoveHouse).Methods("DELETE")

	http.ListenAndServe(":12345", router)
}
