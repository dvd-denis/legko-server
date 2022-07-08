package models

type Step struct {
	Id        int    `db:"id" json:"id"`
	ArticleId int    `db:"article_id" json:"article_id"`
	Title     string `db:"title" json:"title"`
	Content   string `db:"content" json:"content"`
	Num       int    `db:"num" json:"-"`
}

type Image struct {
	Id        int    `db:"id" json:"id"`
	ImageName string `db:"image_name" json:"image_name"`
	Image     string `db:"image" json:"image"`
	StepId    int    `db:"step_id" json:"step_id"`
}
