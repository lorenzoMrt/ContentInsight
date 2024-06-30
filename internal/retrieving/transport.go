package retrieving

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	kitlog "github.com/go-kit/log"
	cr "github.com/lorenzoMrt/ContentInsight/internal"
	"github.com/lorenzoMrt/ContentInsight/kit/query"
)

func MakeHandler(b query.Bus, logger kitlog.Logger) gin.HandlerFunc {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

	getContentHandler := kithttp.NewServer(
		makeGetByUUIDEndpoint(b),
		decodeGetContentRequest,
		encodeResponse,
		opts...,
	)

	return func(ctx *gin.Context) {
		getContentHandler.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func decodeGetContentRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getContentByUUIDRequest
	request.UUID = r.URL.Query().Get("uuid")

	return request, nil
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
	case errors.Is(err, cr.ErrContentNotFound):
		return http.StatusNotFound
	case errors.Is(err, cr.ErrInvalidContentID):
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
