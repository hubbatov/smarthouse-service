package restapi

//RESTUser (data)
type RESTUser struct {
	Name     string `gorm:"type:varchar(100)" json:"name"`
	Login    string `gorm:"type:varchar(100);unique" json:"login"`
	Password string `gorm:"type:varchar(100)" json:"password"`
}
