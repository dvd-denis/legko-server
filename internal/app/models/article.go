package models

type Article struct {
	Id       int    `db:"id"`
	Title    string `db:"title"`
	IconName string `db:"icon_name"`
	Icon     []byte `db:"icon"`
	Url      string `db:"url"`
	Color    string `db:"color"`
}
