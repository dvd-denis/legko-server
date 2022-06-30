package models

type Article struct {
	Id       int
	Title    string
	IconName string `db:"icon_name"`
	Icon     []byte
	Color    string
}
