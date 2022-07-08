package models

type Article struct {
	Id      int    `db:"id" json:"id"`
	Title   string `db:"title" json:"title"`
	GroupId int    `db:"group_id" json:"group_id"`
}
