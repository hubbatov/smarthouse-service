package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"anybodyhere/restapi"

	"github.com/gorilla/mux"
)

// Service for data access
type Service struct {
	dataManager DatabaseManager
}

//Run service
func (s *Service) Run() {
	log.Print("Starting service ...")

	s.dataManager.createDb()

	router := mux.NewRouter()

	router.HandleFunc("/users", s.getUsers).Methods("GET")
	router.HandleFunc("/users/register", s.registerUser).Methods("POST")

	http.ListenAndServe(":12345", router)
}

//Users

func (s *Service) getUsers(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(s.dataManager.users())
}

func (s *Service) registerUser(w http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")

	if contentType != "application/json" {
		handleError(GenerateCustomError("Content-Type is not application/json"))
	}

	body, err := ioutil.ReadAll(req.Body)

	handleError(ConvertCustomError(err))

	var userdata restapi.RESTUser
	err = json.Unmarshal(body, &userdata)

	handleError(ConvertCustomError(err))

	s.dataManager.createUser(userdata)
}
