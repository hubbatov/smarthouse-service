package restapi

//RESTSensor represents REST data for sensors
type RESTSensor struct {
	Name    string `gorm:"type:varchar(100)" json:"name"`
	Tag     string `gorm:"type:varchar(100);unique" json:"tag"`
	HouseID int    `json:"-"`
}

//RESTSensorDataFilter represents filter for sensor data
type RESTSensorDataFilter struct {
	After string `json:"after"`
}

//RESTSensorData RESTSensorData prepresents REST data for sensor measurings
type RESTSensorData struct {
	SensorID int    `json:"-"`
	Data     string `json:"data"`
}
