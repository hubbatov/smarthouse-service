package models

import "anybodyhere/restapi"

//Sensor represents sensor in house
type Sensor struct {
	ID int `gorm:"primary_key" json:"id"`
	restapi.RESTSensor
}

//CreateSensor creates new sensor for house
func CreateSensor(sensordata restapi.RESTSensor) Sensor {
	a := Sensor{}
	a.HouseID = sensordata.HouseID
	a.Name = sensordata.Name
	return a
}

//TableName for Sensors
func (Sensor) TableName() string {
	return "public.sensors"
}
