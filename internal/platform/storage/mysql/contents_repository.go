package mysql

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/huandu/go-sqlbuilder"
	cr "github.com/lorenzoMrt/ContentInsight/internal"
)

// ContentRepository is a MySQL cr.ContentRepository implementation.
type ContentRepository struct {
	db        *sql.DB
	dbTimeout time.Duration
}

func NewContentRepository(db *sql.DB, dbTimeout time.Duration) *ContentRepository {
	return &ContentRepository{
		db:        db,
		dbTimeout: dbTimeout,
	}
}

func marshalField(data interface{}) ([]byte, error) {
	json, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return json, nil
}

func unmarshalField(data []byte, v interface{}) error {
	err := json.Unmarshal(data, v)
	if err != nil {
		return err
	}
	return nil
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

	ctxTimeout, cancel := context.WithTimeout(context.Background(), repo.dbTimeout)
	defer cancel()

	_, err = repo.db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		if err == context.Canceled {
			return fmt.Errorf("query canceled due to context cancellation")

		} else {

			return fmt.Errorf("error trying to persist content on database: %v", err)
		}
	}

	return nil
}

func (repo *ContentRepository) QueryByUuid(ctx context.Context, uuid string) (cr.Content, error) {
	contentSqlStruct := sqlbuilder.NewStruct(new(sqlContent))
	query, args := contentSqlStruct.SelectFrom(sqlContentTable).Where("uuid = ?", uuid).Build()
	ctxTimeout, cancel := context.WithTimeout(context.Background(), repo.dbTimeout)
	defer cancel()

	row := repo.db.QueryRowContext(ctxTimeout, query, args)
	var sqlContent sqlContent
	if err := row.Scan(contentSqlStruct.Addr(&sqlContent)...); err != nil {
		if err == sql.ErrNoRows {
			return cr.Content{}, nil
		}
		return cr.Content{}, err
	}

	var categories []string
	if err := unmarshalField(sqlContent.Categories, &categories); err != nil {
		return cr.Content{}, err
	}
	var tags []string
	if err := unmarshalField(sqlContent.Tags, &tags); err != nil {
		return cr.Content{}, err
	}
	var metadata cr.Metadata
	if err := unmarshalField(sqlContent.Metadata, &metadata); err != nil {
		return cr.Content{}, err
	}
	content, err := cr.NewContent(
		sqlContent.Uuid,
		sqlContent.Title,
		sqlContent.Description,
		sqlContent.ContentType,
		categories,
		tags,
		sqlContent.Author,
		sqlContent.PublicationDate,
		sqlContent.ContentURL,
		sqlContent.Duration,
		sqlContent.Language,
		sqlContent.CoverImage,
		metadata,
		sqlContent.Status,
		sqlContent.Source,
		sqlContent.Visibility)

	if err != nil {
		return cr.Content{}, err
	}

	return content, nil
}
