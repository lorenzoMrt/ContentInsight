package courses

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.POST("/create", CreateHandler())

	// Prepare the request payload
	contentReq := ContentRequest{
		Uuid:            uuid.New(),
		Title:           "Introducción a Microservicios",
		Description:     "Un artículo completo sobre los fundamentos de los microservicios.",
		ContentType:     "article",
		Categories:      []string{"Tecnología", "Desarrollo de Software"},
		Tags:            []string{"microservicios", "arquitectura", "desarrollo"},
		Author:          "Juan Pérez",
		PublicationDate: time.Date(2024, 6, 1, 12, 0, 0, 0, time.UTC),
		ContentURL:      "https://ejemplo.com/introduccion-a-microservicios",
		Duration:        nil,
		Language:        "es",
		CoverImage:      "https://ejemplo.com/imagenes/introduccion-a-microservicios.jpg",
		Metadata: Metadata{
			Views:    1500,
			Likes:    200,
			Comments: 10,
		},
		Status:     "publicado",
		Source:     "Blog Ejemplo",
		Visibility: "publico",
	}

	jsonData, err := json.Marshal(contentReq)
	assert.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, "/create", bytes.NewBuffer(jsonData))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder to capture the response
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response ContentRequest
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, contentReq.Uuid, response.Uuid)
	assert.Equal(t, contentReq.Title, response.Title)
	assert.Equal(t, contentReq.Description, response.Description)
	assert.Equal(t, contentReq.ContentType, response.ContentType)
	assert.Equal(t, contentReq.Categories, response.Categories)
	assert.Equal(t, contentReq.Tags, response.Tags)
	assert.Equal(t, contentReq.Author, response.Author)
	assert.Equal(t, contentReq.PublicationDate, response.PublicationDate)
	assert.Equal(t, contentReq.ContentURL, response.ContentURL)
	assert.Equal(t, contentReq.Duration, response.Duration)
	assert.Equal(t, contentReq.Language, response.Language)
	assert.Equal(t, contentReq.CoverImage, response.CoverImage)
	assert.Equal(t, contentReq.Metadata, response.Metadata)
	assert.Equal(t, contentReq.Status, response.Status)
	assert.Equal(t, contentReq.Source, response.Source)
	assert.Equal(t, contentReq.Visibility, response.Visibility)
}
