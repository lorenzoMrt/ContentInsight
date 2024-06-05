package contents

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lorenzoMrt/ContentInsight/internal/platform/storage/storagemocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateHandler(t *testing.T) {
	contentRepository := new(storagemocks.ContentRepository)
	contentRepository.On("Save", mock.Anything, mock.AnythingOfType("cr.Content")).Return(nil)
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.POST("/create", CreateHandler(contentRepository))

	t.Run("200 OK", func(t *testing.T) {
		// Prepare the request payload
		contentReq := ContentRequest{
			Uuid:            uuid.New().String(),
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

		contentRepository.AssertCalled(t, "Save", mock.Anything, mock.AnythingOfType("cr.Content"))

	})

	t.Run("400 BAD REQUEST", func(t *testing.T) {
		contentReq := ContentRequest{
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

		res := w.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)

	})
}
