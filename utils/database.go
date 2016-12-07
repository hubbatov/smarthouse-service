package utils

import (
	"fmt"

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

//DatabaseManager (holds database connection)
type DatabaseManager struct {
	dataBase *gorm.DB
}

func (d *DatabaseManager) createDb() {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		DbHost, DbPort, DbUser, DbPassword, DbName)
	db, err := gorm.Open("postgres", dbinfo)
	handleError(ConvertCustomError(err))

	db.AutoMigrate(&models.User{})

	d.dataBase = db
}

//Users

func (d *DatabaseManager) users() []models.User {
	var table []models.User
	d.dataBase.Order("id").Limit(10).Find(&table)
	return table
}

func (d *DatabaseManager) createUser(userdata restapi.RESTUser) {
	u := models.CreateUser(userdata)
	d.dataBase.Create(&u)
}
