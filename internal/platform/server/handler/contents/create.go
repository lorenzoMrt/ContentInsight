package contents

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ContentRequest struct {
	Uuid            uuid.UUID `json:"uuid"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	ContentType     string    `json:"contentType"`
	Categories      []string  `json:"categories"`
	Tags            []string  `json:"tags"`
	Author          string    `json:"author"`
	PublicationDate time.Time `json:"publicationDate"`
	ContentURL      string    `json:"contentUrl"`
	Duration        *int      `json:"duration"`
	Language        string    `json:"language"`
	CoverImage      string    `json:"coverImage"`
	Metadata        Metadata  `json:"metadata"`
	Status          string    `json:"status"`
	Source          string    `json:"source"`
	Visibility      string    `json:"visibility"`
}
type Metadata struct {
	Views    int `json:"views"`
	Likes    int `json:"likes"`
	Comments int `json:"comments"`
}

// CreateHandler returns an HTTP handler for courses creation.
func CreateHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req ContentRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		jsonData, err := json.Marshal(req)
		if err != nil {
			fmt.Println("Error serializing to JSON:", err)
			return
		}

		// Convert the JSON byte slice to a string
		jsonString := string(jsonData)

		ctx.String(http.StatusCreated, jsonString)
	}
}
