package models

type Article struct {
	Id       int    `db:"id" json:"id"`
	Title    string `db:"title" json:"title"`
	IconName string `db:"icon_name" json:"icon_name"`
	Icon     string `db:"icon" json:"icon"`
	Url      string `db:"url" json:"url"`
	Color    string `db:"color" json:"color"`
	Wifi     bool   `db:"wifi" json:"wifi"`
	Question bool   `db:"question" json:"-"`
}
