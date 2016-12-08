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

	DBManager.dataBase = db
}

func (d *DatabaseManager) users() []models.User {
	var table []models.User
	d.dataBase.Order("id").Limit(10).Find(&table)
	return table
}

func (d *DatabaseManager) createUser(userdata restapi.RESTUser) []error {
	u := models.CreateUser(userdata)
	return d.dataBase.Create(&u).GetErrors()
}
