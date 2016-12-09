package restapi

//RESTHouse represents REST data for houses
type RESTHouse struct {
	UserID  int    `gorm:"index" json:"-"`
	Name    string `gorm:"type:varchar(100)" json:"name"`
	Address string `gorm:"type:varchar(100)" json:"address"`
}
