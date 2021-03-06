package controllers

import (
	"fmt"

	"smarthouse-service/errors"
	"smarthouse-service/models"
	"smarthouse-service/restapi"

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
	DbName = "smarthouse"
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
	db.AutoMigrate(&models.Command{})

	DBManager.dataBase = db
}

func (d *DatabaseManager) users() []models.User {
	var table []models.User
	d.dataBase.Order("id").Find(&table)
	return table
}

func (d *DatabaseManager) user(userLogin, userPassword string) []models.User {
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

func (d *DatabaseManager) sensorID(sensorTag string) int {
	var sensor models.Sensor

	errors := d.dataBase.Where("tag = ?", sensorTag).First(&sensor).GetErrors()
	if len(errors) > 0 {
		return -1
	}

	return sensor.ID
}

func (d *DatabaseManager) userIDFromSensor(sensorID int) int {
	var sensor models.Sensor
	errors := d.dataBase.Where("id = ?", sensorID).First(&sensor).GetErrors()
	if len(errors) > 0 {
		return -1
	}

	var house models.House
	errors = d.dataBase.Where("id = ?", sensor.HouseID).First(&house).GetErrors()
	if len(errors) > 0 {
		return -1
	}

	var user models.User
	errors = d.dataBase.Where("id = ?", house.UserID).First(&user).GetErrors()
	if len(errors) > 0 {
		return -1
	}

	return user.ID
}

func (d *DatabaseManager) addSensor(sensordata restapi.RESTSensor) []error {
	h := models.CreateSensor(sensordata)
	return d.dataBase.Create(&h).GetErrors()
}

func (d *DatabaseManager) editSensor(sensorID int, sensordata restapi.RESTSensor) []error {
	errors := d.dataBase.Model(&models.Sensor{}).Where("id = ?", sensorID).Update("name", sensordata.Name).GetErrors()
	if len(errors) > 0 {
		return errors
	}

	return d.dataBase.Model(&models.Sensor{}).Where("id = ?", sensorID).Update("tag", sensordata.Tag).GetErrors()
}

func (d *DatabaseManager) removeSensor(sensorID int) []error {
	return d.dataBase.Where("id = ?", sensorID).Delete(&models.Sensor{}).GetErrors()
}

func (d *DatabaseManager) commands(houseID int) []models.Command {
	var table []models.Command
	d.dataBase.Where("house_id = ?", houseID).Order("id").Find(&table)
	return table
}

func (d *DatabaseManager) userIDFromCommand(commandID int) int {
	var command models.Command
	errors := d.dataBase.Where("id = ?", commandID).First(&command).GetErrors()
	if len(errors) > 0 {
		return -1
	}

	var house models.House
	errors = d.dataBase.Where("id = ?", command.HouseID).First(&house).GetErrors()
	if len(errors) > 0 {
		return -1
	}

	var user models.User
	errors = d.dataBase.Where("id = ?", house.UserID).First(&user).GetErrors()
	if len(errors) > 0 {
		return -1
	}

	return user.ID
}

func (d *DatabaseManager) addCommand(command restapi.RESTCommand) []error {
	h := models.CreateCommand(command)
	return d.dataBase.Create(&h).GetErrors()
}

func (d *DatabaseManager) editCommand(commandID int, command restapi.RESTCommand) []error {
	errors := d.dataBase.Model(&models.Command{}).Where("id = ?", commandID).Update("name", command.Name).GetErrors()
	if len(errors) > 0 {
		return errors
	}

	errors = d.dataBase.Model(&models.Command{}).Where("id = ?", commandID).Update("query", command.Query).GetErrors()
	if len(errors) > 0 {
		return errors
	}

	errors = d.dataBase.Model(&models.Command{}).Where("id = ?", commandID).Update("command_type", command.CommandType).GetErrors()
	if len(errors) > 0 {
		return errors
	}

	return d.dataBase.Model(&models.Command{}).Where("id = ?", commandID).Update("available_values", command.AvailableValues).GetErrors()
}

func (d *DatabaseManager) removeCommand(commandID int) []error {
	return d.dataBase.Where("id = ?", commandID).Delete(&models.Command{}).GetErrors()
}

func (d *DatabaseManager) sensordata(sensorID int, after string, before string) []models.SensorData {
	var table []models.SensorData

	if len(after) != 0 && len(before) != 0 {
		d.dataBase.Where("sensor_id = ? AND time > ? AND time <= ?", sensorID, after, before).Order("time").Find(&table)
	} else if len(after) == 0 {
		d.dataBase.Where("sensor_id = ? AND time <= ?", sensorID, before).Order("time").Find(&table)
	} else if len(before) == 0 {
		d.dataBase.Where("sensor_id = ? AND time > ?", sensorID, after).Order("time").Find(&table)
	} else {
		d.dataBase.Where("sensor_id = ?", sensorID).Order("time").Find(&table)
	}

	return table
}

func (d *DatabaseManager) addSensorData(sensordata restapi.RESTSensorData) []error {
	sd := models.CreateSensorData(sensordata)
	return d.dataBase.Create(&sd).GetErrors()
}
