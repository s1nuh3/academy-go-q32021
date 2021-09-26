package models

//Users -- Model to work with an external API later on
type Users struct {
	ID     int    `csv,json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Gender string `json:"gender"`
	Status bool   `json:"status"`
}
