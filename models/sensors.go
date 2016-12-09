package models

import (
	"anybodyhere/restapi"
	"time"
)

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

//SensorData represents sensor record on time
type SensorData struct {
	ID int `gorm:"primary_key" json:"-"`
	restapi.RESTSensorData
	Time time.Time `json:"time"`
}

//CreateSensorData creates new sensor record
func CreateSensorData(sensordata restapi.RESTSensorData) SensorData {
	a := SensorData{}
	a.SensorID = sensordata.SensorID
	a.Data = sensordata.Data
	a.Time = time.Now()
	return a
}

//TableName for SensorDatas
func (SensorData) TableName() string {
	return "public.sensordatas"
}
