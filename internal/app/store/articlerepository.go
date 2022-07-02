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

func (r *ArticleRepository) GetSteps(id int) ([]models.Step, error) {
	var steps []models.Step

	query := fmt.Sprintf("SELECT * FROM %s WHERE article_id = ?", step_tabel)

	err := r.store.db.Select(&steps, query, id)

	return steps, err
}
