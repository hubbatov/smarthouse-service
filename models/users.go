package models

import (
	"anybodyhere/restapi"
	"time"
)

//User (typical user of service)
type User struct {
	ID    uint      `gorm:"primary_key" json:"-"`
	Since time.Time `json:"since"`
	restapi.RESTUser
}

//CreateUser (creating new user)
func CreateUser(userdata restapi.RESTUser) User {
	a := User{}
	a.Name = userdata.Name
	a.Login = userdata.Name
	a.Password = userdata.Name
	a.Since = time.Now()
	return a
}

//TableName for Users
func (User) TableName() string {
	return "public.users"
}
