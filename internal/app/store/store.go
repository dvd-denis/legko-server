package store

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	article_table = "articles"
	step_table    = "steps"
	image_table   = "images"
)

type Store struct {
	config            *Config
	db                *sqlx.DB
	articleRepository *ArticleRepository
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (s *Store) Open() error {
	db, err := sqlx.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) Article() *ArticleRepository {
	if s.articleRepository != nil {
		return s.articleRepository
	}

	s.articleRepository = &ArticleRepository{
		store: s,
	}

	return s.articleRepository
}
