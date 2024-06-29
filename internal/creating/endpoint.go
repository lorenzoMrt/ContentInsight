package creating

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/lorenzoMrt/ContentInsight/kit/command"
)

type contentRequest struct {
	Uuid            string
	Title           string
	Description     string
	ContentType     string
	Categories      []string
	Tags            []string
	Author          string
	PublicationDate time.Time
	ContentURL      string
	Duration        *int
	Language        string
	CoverImage      string
	Metadata        metadataRequest
	Status          string
	Source          string
	Visibility      string
}
type metadataRequest struct {
	Views    int
	Likes    int
	Comments int
}

type createContentResponse struct {
	Err error `json:"error,omitempty"`
}

// Convert ContentRequest.Metadata to cr.Metadata
func toCreatingMetadata(metadata metadataRequest) Metadata {
	return Metadata(metadata)
}

func (r createContentResponse) error() error { return r.Err }

func makeCreateContentEndpoint(commandBus command.Bus) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(contentRequest)
		contentCommand := NewContentCommand(req.Uuid, req.Title, req.Description, req.ContentType, req.Categories, req.Tags, req.Author, req.PublicationDate, req.ContentURL, req.Duration, req.Language, req.CoverImage, toCreatingMetadata(req.Metadata), req.Status, req.Source, req.Visibility)
		err := commandBus.Dispatch(ctx, contentCommand)
		return createContentResponse{Err: err}, nil
	}
}
