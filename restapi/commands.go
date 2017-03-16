package restapi

//RESTCommand represents REST data for control commands
type RESTCommand struct {
	HouseID         int    `gorm:"index" json:"-"`
	Name            string `gorm:"type:varchar(100)" json:"name"`
	Query           string `gorm:"type:varchar(100)" json:"query"`
	CommandType     string `gorm:"type:varchar(10)" json:"command_type"`
	AvailableValues string `gorm:"type:varchar(100)" json:"available_values"`
}
