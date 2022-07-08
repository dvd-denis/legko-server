package models

type Group struct {
	Id       int    `db:"id" json:"id"`
	Title    string `db:"title" json:"title"`
	IconName string `db:"icon_name" json:"icon_name"`
	Icon     string `db:"icon" json:"icon"`
	Color    string `db:"color" json:"color"`
	Wifi     bool   `db:"wifi" json:"wifi"`
	Model    string `db:"model" json:"model"`
}
