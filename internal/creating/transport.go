package creating

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	kitlog "github.com/go-kit/log"
	cr "github.com/lorenzoMrt/ContentInsight/internal"
	"github.com/lorenzoMrt/ContentInsight/kit/command"
)

type metadataBody struct {
	Views    int `json:"views"`
	Likes    int `json:"likes"`
	Comments int `json:"comments"`
}

var body struct {
	Uuid            string       `json:"uuid" binding:"required"`
	Title           string       `json:"title"`
	Description     string       `json:"description"`
	ContentType     string       `json:"contentType"`
	Categories      []string     `json:"categories"`
	Tags            []string     `json:"tags"`
	Author          string       `json:"author"`
	PublicationDate time.Time    `json:"publicationDate"`
	ContentURL      string       `json:"contentUrl"`
	Duration        *int         `json:"duration"`
	Language        string       `json:"language"`
	CoverImage      string       `json:"coverImage"`
	Metadata        metadataBody `json:"metadata"`
	Status          string       `json:"status"`
	Source          string       `json:"source"`
	Visibility      string       `json:"visibility"`
}

func MakeHandler(b command.Bus, logger kitlog.Logger) gin.HandlerFunc {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

	createContentHandler := kithttp.NewServer(
		makeCreateContentEndpoint(b),
		decodeCreateContentRequest,
		encodeResponse,
		opts...,
	)

	return func(ctx *gin.Context) {
		createContentHandler.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func decodeCreateContentRequest(_ context.Context, r *http.Request) (interface{}, error) {

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return contentRequest{
		Uuid:            body.Uuid,
		Title:           body.Title,
		Description:     body.Description,
		ContentType:     body.ContentType,
		Categories:      body.Categories,
		Tags:            body.Tags,
		Author:          body.Author,
		PublicationDate: body.PublicationDate,
		ContentURL:      body.ContentURL,
		Duration:        body.Duration,
		Language:        body.Language,
		CoverImage:      body.CoverImage,
		Metadata:        toMetadataRequest(body.Metadata),
		Status:          body.Status,
		Source:          body.Source,
		Visibility:      body.Visibility,
	}, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))

	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch {
	case errors.Is(err, cr.ErrInvalidContentID),
		errors.Is(err, cr.ErrEmptyContentTitle):
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

func toMetadataRequest(metadata metadataBody) metadataRequest {
	return metadataRequest(metadata)
}
