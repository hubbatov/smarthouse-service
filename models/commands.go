package models

import "smarthouse-service/restapi"

//Command represents control command in house
type Command struct {
	ID int `gorm:"primary_key" json:"id"`
	restapi.RESTCommand
}

//CreateCommand creates new control command for house
func CreateCommand(command restapi.RESTCommand) Command {
	a := Command{}
	a.HouseID = command.HouseID
	a.Name = command.Name
	a.Query = command.Query
	a.CommandType = command.CommandType
	a.AvailableValues = command.AvailableValues
	return a
}

//TableName for Command
func (Command) TableName() string {
	return "public.commands"
}
