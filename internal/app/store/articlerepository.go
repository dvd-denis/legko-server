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
