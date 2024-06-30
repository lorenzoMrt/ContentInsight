package retrieving

import (
	"context"

	cr "github.com/lorenzoMrt/ContentInsight/internal"
	"github.com/lorenzoMrt/ContentInsight/kit/event"
)

type Service interface {
	QueryContent(ctx context.Context, uuid string) (cr.Content, error)
}

type service struct {
	contentRepository cr.ContentRepository
	eventBus          event.Bus
}

func NewService(contentRepository cr.ContentRepository, eventBus event.Bus) *service {
	return &service{
		contentRepository: contentRepository,
		eventBus:          eventBus,
	}
}

func (s *service) QueryContent(ctx context.Context, uuid string) (cr.Content, error) {
	content, err := s.contentRepository.QueryByUuid(ctx, uuid)
	if err != nil {
		return cr.Content{}, err
	}

	eventBusError := s.eventBus.Publish(ctx, content.PullEvents())
	if eventBusError != nil {
		return content, eventBusError
	}

	return content, nil
}
