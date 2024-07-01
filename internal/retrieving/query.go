package retrieving

import (
	"context"
	"errors"

	cr "github.com/lorenzoMrt/ContentInsight/internal"
	"github.com/lorenzoMrt/ContentInsight/kit/query"
)

const ContentQueryType query.Type = "query.retrieving.Content"

type ContentQuery struct {
	id string
}

func NewContentQuery(contentId cr.ContentID) ContentQuery {
	return ContentQuery{
		id: contentId.String(),
	}
}

func (c ContentQuery) Type() query.Type {
	return ContentQueryType
}

type ContentQueryHandler struct {
	service Service
}

func NewContentQueryHandler(service Service) ContentQueryHandler {
	return ContentQueryHandler{
		service: service,
	}
}

func (h ContentQueryHandler) Handle(ctx context.Context, query query.Query) (interface{}, error) {
	queryContentCmd, ok := query.(ContentQuery)
	if !ok {
		return cr.ContentResponse{}, errors.New("unexpected query")
	}

	queryContent, err := h.service.QueryContent(ctx, queryContentCmd.id)
	if err != nil {
		return cr.ContentResponse{}, err
	}
	return cr.NewContentResponse(queryContent), nil
}
