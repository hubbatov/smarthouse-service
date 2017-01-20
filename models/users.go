package models

import (
	"smarthouse-service/restapi"
	"time"
)

//User represents typical user of service
type User struct {
	ID             int       `gorm:"primary_key" json:"id"`
	Since          time.Time `json:"since"`
	HashedPassword []byte    `json:"-"`
	restapi.RESTUser
	Houses []House `json:"-"`
}

//CreateUser creates new user
func CreateUser(userdata restapi.RESTUser) User {
	a := User{}
	a.Name = userdata.Name
	a.Login = userdata.Login
	a.Password = userdata.Password
	a.Since = time.Now()
	return a
}

//TableName for Users
func (User) TableName() string {
	return "public.users"
}
