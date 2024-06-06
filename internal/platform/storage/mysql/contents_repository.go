package mysql

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/huandu/go-sqlbuilder"
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

func marshalField(data interface{}) ([]byte, error) {
	json, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return json, nil
}

func (repo *ContentRepository) Save(ctx context.Context, content cr.Content) error {

	contentSqlStruct := sqlbuilder.NewStruct(new(sqlContent))
	jsonCategories, err := marshalField(content.Categories())
	if err != nil {
		return err
	}

	jsonTags, err := marshalField(content.Tags())
	if err != nil {
		return err
	}

	jsonMetadata, err := marshalField(content.Metadata())
	if err != nil {
		return err
	}

	query, args := contentSqlStruct.InsertInto(sqlContentTable, sqlContent{
		Uuid:            content.ID().String(),
		Title:           content.Title().String(),
		Description:     content.Description(),
		ContentType:     content.ContentType(),
		Categories:      []byte(jsonCategories),
		Tags:            []byte(jsonTags),
		Author:          content.Author(),
		PublicationDate: content.PublicationDate(),
		ContentURL:      content.ContentURL(),
		Duration:        content.Duration(),
		Language:        content.Language(),
		CoverImage:      content.CoverImage(),
		Metadata:        []byte(jsonMetadata),
		Status:          content.Status(),
		Source:          content.Source(),
		Visibility:      content.Visibility(),
	}).Build()

	_, err = repo.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist content on database: %v", err)
	}

	return nil
}
