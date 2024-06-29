package health

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type HealthRequest struct{}

type HealthResponse struct {
	Status string `json:"status"`
}

func makeHealthEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return HealthResponse{Status: svc.Health()}, nil
	}
}
