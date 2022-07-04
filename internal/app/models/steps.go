package models

type Step struct {
	Id        int    `db:"id"`
	ArticleId int    `db:"article_id"`
	Title     string `db:"title"`
	Content   string `db:"content"`
	Num       int    `db:"num"`
	Wifi      string `db:"wifi"`
}

type Image struct {
	Id        int
	ImageName string `db:"image_name"`
	Image     string
	StepId    int `db:"step_id"`
}
