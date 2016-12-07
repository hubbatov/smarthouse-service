package restapi

//RESTUser (data)
type RESTUser struct {
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
