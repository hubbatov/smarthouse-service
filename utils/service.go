package utils

import (
	"log"
	"net/http"

	"anybodyhere/controllers"

	"github.com/gorilla/mux"
)

// Service for data access
type Service struct {
	usersController controllers.UsersController
}

//Run service
func (s *Service) Run() {
	log.Print("Starting service ...")

	controllers.CreateDb()

	router := mux.NewRouter()

	router.HandleFunc("/users", s.usersController.GetUsers).Methods("GET")
	router.HandleFunc("/users", s.usersController.RegisterUser).Methods("POST")

	router.HandleFunc("/login", s.usersController.Login).Methods("POST")

	http.ListenAndServe(":12345", router)
}
