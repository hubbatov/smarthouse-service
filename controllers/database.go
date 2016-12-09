package controllers

import (
	"fmt"

	"anybodyhere/errors"
	"anybodyhere/models"
	"anybodyhere/restapi"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //pg adapter
)

const (
	//DbHost - database host
	DbHost = "localhost"
	//DbPort - database port
	DbPort = 5432
	//DbUser - database user
	DbUser = "postgres"
	//DbPassword - database password
	DbPassword = "123456"
	//DbName - database db name
	DbName = "anybodyhere"
)

//DBManager is a standalone DatabaseManager object
var DBManager DatabaseManager

//DatabaseManager (holds database connection)
type DatabaseManager struct {
	dataBase *gorm.DB
}

// CreateDb creates new DatabaseManager
func CreateDb() {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		DbHost, DbPort, DbUser, DbPassword, DbName)
	db, err := gorm.Open("postgres", dbinfo)

	errors.HandleError(errors.ConvertCustomError(err))

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.House{})
	db.AutoMigrate(&models.Sensor{})
	db.AutoMigrate(&models.SensorData{})

	DBManager.dataBase = db
}

func (d *DatabaseManager) users() []models.User {
	var table []models.User
	d.dataBase.Order("id").Find(&table)
	return table
}

func (d *DatabaseManager) createUser(userdata restapi.RESTUser) []error {
	u := models.CreateUser(userdata)
	return d.dataBase.Create(&u).GetErrors()
}

func (d *DatabaseManager) houses(userID int) []models.House {
	var table []models.House
	d.dataBase.Where("user_id = ?", userID).Order("id").Find(&table)
	return table
}

func (d *DatabaseManager) addHouse(housedata restapi.RESTHouse) []error {
	h := models.CreateHouse(housedata)
	return d.dataBase.Create(&h).GetErrors()
}

func (d *DatabaseManager) editHouse(houseID int, housedata restapi.RESTHouse) []error {
	errors := d.dataBase.Model(&models.House{}).Where("id = ?", houseID).Update("name", housedata.Name).GetErrors()
	if len(errors) > 0 {
		return errors
	}
	return d.dataBase.Model(&models.House{}).Where("id = ?", houseID).Update("address", housedata.Address).GetErrors()
}

func (d *DatabaseManager) removeHouse(houseID int) []error {
	return d.dataBase.Where("id = ?", houseID).Delete(&models.House{}).GetErrors()
}

func (d *DatabaseManager) sensors(houseID int) []models.Sensor {
	var table []models.Sensor
	d.dataBase.Where("house_id = ?", houseID).Order("id").Find(&table)
	return table
}

func (d *DatabaseManager) addSensor(sensordata restapi.RESTSensor) []error {
	h := models.CreateSensor(sensordata)
	return d.dataBase.Create(&h).GetErrors()
}

func (d *DatabaseManager) editSensor(sensorID int, sensordata restapi.RESTSensor) []error {
	return d.dataBase.Model(&models.Sensor{}).Where("id = ?", sensorID).Update("name", sensordata.Name).GetErrors()
}

func (d *DatabaseManager) removeSensor(sensorID int) []error {
	return d.dataBase.Where("id = ?", sensorID).Delete(&models.Sensor{}).GetErrors()
}

func (d *DatabaseManager) sensordata(sensorID int) []models.SensorData {
	var table []models.SensorData
	d.dataBase.Where("sensor_id = ?", sensorID).Order("time").Find(&table)
	return table
}

func (d *DatabaseManager) addSensorData(sensordata restapi.RESTSensorData) []error {
	sd := models.CreateSensorData(sensordata)
	return d.dataBase.Create(&sd).GetErrors()
}
