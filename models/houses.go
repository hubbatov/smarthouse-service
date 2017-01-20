package models

import "smarthouse-service/restapi"

//House represents house with sensors
type House struct {
	ID int `gorm:"primary_key" json:"id"`
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
