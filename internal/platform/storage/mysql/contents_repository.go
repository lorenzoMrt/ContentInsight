package mysql

import (
	"context"
	"database/sql"

	cr "github.com/lorenzoMrt/ContentInsight/internal"
)

// ContentRepository is a MySQL cr.ContentRepository implementation.
type ContentRepository struct {
	db *sql.DB
}

func NewContentRepository(db *sql.DB) *ContentRepository {
	return &ContentRepository{
		db: db,
	}
}

func (repo *ContentRepository) Save(ctx context.Context, content cr.Content) error {
	//TODO
	return nil
}
