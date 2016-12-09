package models

import "anybodyhere/restapi"

//House represents house with rooms
type House struct {
	ID int `gorm:"primary_key" json:"-"`
	restapi.RESTHouse
}

//CreateHouse creates new house for user
func CreateHouse(housedata restapi.RESTHouse) House {
	a := House{}
	a.UserID = housedata.UserID
	a.Name = housedata.Name
	a.Address = housedata.Address
	return a
}

//TableName for Houses
func (House) TableName() string {
	return "public.houses"
}
