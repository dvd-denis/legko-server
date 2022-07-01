package models

type Step struct {
	Id        int
	ArticleId int `db:"article_id"`
	Title     string
	Content   string
	Num       int
	Wifi      string
}

type Image struct {
	Id        int
	ImageName string `db:"image_name"`
	Image     []byte
	StepId    int `db:"step_id"`
}
