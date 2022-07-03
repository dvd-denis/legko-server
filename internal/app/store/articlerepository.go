package store

import (
	"fmt"

	"github.com/dvd-denis/legko-server/internal/app/models"
)

// ArticleRepository ...
type ArticleRepository struct {
	store *Store
}

// Get all articles
func (r *ArticleRepository) All() ([]models.Article, error) {

	var articles []models.Article

	query := fmt.Sprintf("SELECT * FROM %s", article_table)
	err := r.store.db.Select(&articles, query)

	return articles, err
}

func (r *ArticleRepository) CreateStep(step models.Step) (int, error) {
	var id int
	CreateStepsQuery := fmt.Sprintf("INSERT INTO %s (article_id, title, content, num, wifi) VALUES ($1, $2, $3, $4, $5) RETURNING id", step_tabel)

	row := r.store.db.QueryRow(CreateStepsQuery, step.ArticleId, step.Title, step.Content, step.Num, step.Wifi)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *ArticleRepository) CreateImages(images []models.Image) error {
	CreateImagesQuery := fmt.Sprintf("INSERT INTO %s (step_id, image_name, image) VALUES ($1, $2, $3)", image_tabel)

	for _, image := range images {
		_, err := r.store.db.Exec(CreateImagesQuery, image.StepId, image.ImageName, image.Image)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *ArticleRepository) CreateArticle(article models.Article) (int, error) {
	var id int
	createArticleQuery := fmt.Sprintf("INSERT INTO %s (title, icon_name, icon, url, color) VALUES ($1, $2, $3, $4, $5) RETURNING id", article_table)

	row := r.store.db.QueryRow(createArticleQuery, article.Title, article.IconName, article.Icon, article.Url, article.Color)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *ArticleRepository) GetSteps(id int) ([]models.Step, error) {
	var steps []models.Step

	query := fmt.Sprintf("SELECT * FROM %s WHERE article_id = ?", step_tabel)

	err := r.store.db.Select(&steps, query, id)

	return steps, err
}
