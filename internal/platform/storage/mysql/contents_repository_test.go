package mysql

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	cr "github.com/lorenzoMrt/ContentInsight/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Produce_error(t *testing.T) {
	// Arrange
	// Initialize your variables and mocks here
	contentID, contentTitle, contentDescription, contentType, author := "ae6907fd-4dcf-4df7-9dbb-916c697c2405", "Title", "Un artículo completo sobre los fundamentos de los microservicios.", "article", "Juan peres"
	categories, tags := []string{"Tecnología", "Desarrollo de Software"}, []string{"microservicios", "arquitectura", "desarrollo"}
	jsonCategories, err := json.Marshal(categories)
	require.NoError(t, err)
	jsonTags, err := json.Marshal(tags)
	require.NoError(t, err)
	jsonMetadata, err := json.Marshal(cr.Metadata{})
	require.NoError(t, err)
	publicationDate := time.Date(2024, 6, 1, 12, 0, 0, 0, time.UTC)
	url := "https://ejemplo.com/introduccion-a-microservicios"

	content, err := cr.NewContent(contentID, contentTitle, contentDescription, contentType, categories, tags, author, publicationDate, url, nil, "video", "nil", cr.Metadata{}, "nil", "nil", "")
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec(
		"INSERT INTO contents (uuid, title, description, contentType, categories, tags, author, publicationDate, contentUrl, duration, language, coverImage, metadata, status, source, visibility) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)").
		WithArgs(contentID, contentTitle, contentDescription, contentType, jsonCategories, jsonTags, author, publicationDate, url, nil, "video", "nil", jsonMetadata, "nil", "nil", "").
		WillReturnError(errors.New("something-failed"))

	repo := NewContentRepository(db, 5*time.Second)
	// Act
	// Call the function you want to test
	err = repo.Save(context.Background(), content)
	// Assert
	// Use assert package for your assertions
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}

func Test_Save_Succeed(t *testing.T) {
	// Arrange
	// Initialize your variables and mocks here
	contentID, contentTitle, contentDescription, contentType, author := "ae6907fd-4dcf-4df7-9dbb-916c697c2405", "Title", "Un artículo completo sobre los fundamentos de los microservicios.", "article", "Juan peres"
	categories, tags := []string{"Tecnología", "Desarrollo de Software"}, []string{"microservicios", "arquitectura", "desarrollo"}
	jsonCategories, err := json.Marshal(categories)
	require.NoError(t, err)
	jsonTags, err := json.Marshal(tags)
	require.NoError(t, err)
	jsonMetadata, err := json.Marshal(cr.Metadata{})
	require.NoError(t, err)
	publicationDate := time.Date(2024, 6, 1, 12, 0, 0, 0, time.UTC)
	url := "https://ejemplo.com/introduccion-a-microservicios"

	content, err := cr.NewContent(contentID, contentTitle, contentDescription, contentType, categories, tags, author, publicationDate, url, nil, "video", "nil", cr.Metadata{}, "nil", "nil", "")
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec(
		"INSERT INTO contents (uuid, title, description, contentType, categories, tags, author, publicationDate, contentUrl, duration, language, coverImage, metadata, status, source, visibility) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)").
		WithArgs(contentID, contentTitle, contentDescription, contentType, jsonCategories, jsonTags, author, publicationDate, url, nil, "video", "nil", jsonMetadata, "nil", "nil", "").
		WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewContentRepository(db, 5*time.Second)
	// Act
	// Call the function you want to test
	err = repo.Save(context.Background(), content)
	// Assert
	// Use assert package for your assertions
	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
}
