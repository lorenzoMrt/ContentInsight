package contents

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	cr "github.com/lorenzoMrt/ContentInsight/internal"
	"github.com/lorenzoMrt/ContentInsight/internal/creating"
	"github.com/lorenzoMrt/ContentInsight/kit/command"
)

type ContentRequest struct {
	Uuid            string    `json:"uuid" binding:"required"`
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

// Convert ContentRequest.Metadata to cr.Metadata
func toCreatingMetadata(metadata Metadata) creating.Metadata {
	return creating.Metadata{
		Views:    metadata.Views,
		Likes:    metadata.Likes,
		Comments: metadata.Comments,
	}
}

// CreateHandler returns an HTTP handler for courses creation.
func CreateHandler(commandBus command.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req ContentRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := commandBus.Dispatch(ctx, creating.NewContentCommand(req.Uuid, req.Title, req.Description, req.ContentType, req.Categories, req.Tags, req.Author, req.PublicationDate, req.ContentURL, req.Duration, req.Language, req.CoverImage, toCreatingMetadata(req.Metadata), req.Status, req.Source, req.Visibility))

		if err != nil {
			switch {
			case errors.Is(err, cr.ErrInvalidContentID),
				errors.Is(err, cr.ErrEmptyContentTitle):
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			default:
				ctx.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		}

		ctx.Status(http.StatusCreated)
	}
}
