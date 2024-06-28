package health

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	kitlog "github.com/go-kit/log"
)

func MakeHandler(svc Service, logger kitlog.Logger) gin.HandlerFunc {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

	createHealtHandler := kithttp.NewServer(
		makeHealthEndpoint(svc),
		decodeHealthRequest,
		encodeResponse,
		opts...,
	)

	return func(ctx *gin.Context) {
		createHealtHandler.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func decodeHealthRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return HealthRequest{}, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
