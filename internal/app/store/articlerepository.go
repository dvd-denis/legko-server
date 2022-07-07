package store

import (
	"fmt"
	"strings"

	"github.com/dvd-denis/legko-server/internal/app/models"
)

// ArticleRepository ...
type ArticleRepository struct {
	store *Store
}

// Get all articles or questions
func (r *ArticleRepository) All(model string, isQuestion bool) ([]models.Article, error) {

	var articles []models.Article

	query := fmt.Sprintf("SELECT * FROM %s WHERE question = $1 AND model = $2", article_table)
	err := r.store.db.Select(&articles, query, isQuestion, model)

	return articles, err
}

func (r *ArticleRepository) CreateStep(step models.Step) (int, error) {
	var id int
	CreateStepsQuery := fmt.Sprintf("INSERT INTO %s (article_id, title, content, num) VALUES ($1, $2, $3, $4) RETURNING id", step_table)
	row := r.store.db.QueryRow(CreateStepsQuery, step.ArticleId, step.Title, step.Content, step.Num)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *ArticleRepository) CreateImage(image models.Image) error {
	CreateImagesQuery := fmt.Sprintf("INSERT INTO %s (step_id, image_name, image) VALUES ($1, $2, $3)", image_table)

	_, err := r.store.db.Exec(CreateImagesQuery, image.StepId, image.ImageName, image.Image)
	if err != nil {
		return err
	}

	return nil
}

func (r *ArticleRepository) CreateArticle(article models.Article) (int, error) {
	var id int
	createArticleQuery := fmt.Sprintf("INSERT INTO %s (title, icon_name, icon, url, color, wifi, question) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id", article_table)

	row := r.store.db.QueryRow(createArticleQuery, article.Title, article.IconName, article.Icon, article.Url, article.Color, article.Wifi, article.Question)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *ArticleRepository) GetSteps(id int) ([]models.Step, error) {
	var steps []models.Step

	query := fmt.Sprintf("SELECT * FROM %s WHERE article_id = $1", step_table)

	rows, err := r.store.db.Queryx(query, id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var step models.Step
		if err := rows.StructScan(&step); err != nil {
			return nil, err
		}

		images, err := r.GetImagesAsStep(step.Id)
		if err != nil {
			return nil, err
		}

		for _, image := range images {
			img := `<img src="data:image/png;base64, ` + image.Image + ` ">`
			step.Content = strings.ReplaceAll(step.Content, "{{"+image.ImageName+"}}", img)
		}

		steps = append(steps, step)
	}

	return steps, err
}

func (r *ArticleRepository) GetImagesAsStep(id int) ([]models.Image, error) {
	var images []models.Image

	query := fmt.Sprintf("Select * from %s where step_id = $1", image_table)

	err := r.store.db.Select(&images, query, id)

	return images, err
}

func (r *ArticleRepository) DeleteArticle(id int) error {
	DeleteQuery := fmt.Sprintf("DELETE FROM %s WHERE id = $1;", article_table)
	_, err := r.store.db.Exec(DeleteQuery, id)
	if err != nil {
		return err
	}

	return nil
}
