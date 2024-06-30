package retrieving

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"
	cr "github.com/lorenzoMrt/ContentInsight/internal"
	"github.com/lorenzoMrt/ContentInsight/kit/query"
)

type getContentByUUIDRequest struct {
	UUID string `json:"uuid"`
}

type metadataBody struct {
	Views    int `json:"views"`
	Likes    int `json:"likes"`
	Comments int `json:"comments"`
}

type content struct {
	Uuid            string       `json:"uuid"`
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

type getByUUIDResponse struct {
	Data  content `json:"data,omitempty"`
	Error string  `json:"error,omitempty"`
}

func makeGetByUUIDEndpoint(queryBus query.Bus) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getContentByUUIDRequest)
		contentId, err := cr.NewContentID(req.UUID)
		if err != nil {
			return getByUUIDResponse{Error: err.Error()}, nil
		}
		contentQuery := NewContentQuery(contentId)
		r, err := queryBus.Ask(ctx, contentQuery)
		if err != nil {
			return getByUUIDResponse{Error: err.Error()}, nil
		}
		data := r.(content)
		return getByUUIDResponse{Data: data}, nil
	}
}
