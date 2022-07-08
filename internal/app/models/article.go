package models

type Article struct {
	Id      int    `db:"id" json:"id"`
	Title   string `db:"title" json:"title"`
	Tags    string `db:"tags" json:"tags"`
	GroupId int    `db:"group_id" json:"group_id"`
}
