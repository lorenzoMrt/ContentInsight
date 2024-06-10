package creating

import (
	"context"
	"time"

	cr "github.com/lorenzoMrt/ContentInsight/internal"
	"github.com/lorenzoMrt/ContentInsight/kit/event"
)

type ContentService struct {
	contentRepository cr.ContentRepository
	eventBus          event.Bus
}

func NewContentService(contentRepository cr.ContentRepository, eventBus event.Bus) ContentService {
	return ContentService{
		contentRepository: contentRepository,
		eventBus:          eventBus,
	}
}

func (s ContentService) CreateContent(ctx context.Context, uuid string, title string, description string, contentType string, categories []string, tags []string, author string, publicationDate time.Time, contentURL string, duration *int, language string, coverImage string, metadata cr.Metadata, status string, source string, visibility string) error {
	content, err := cr.NewContent(uuid, title, description, contentType, categories, tags, author, publicationDate, contentURL, duration, language, coverImage, metadata, status, source, visibility)

	if err != nil {
		return err
	}
	if err := s.contentRepository.Save(ctx, content); err != nil {
		return err
	}

	return s.eventBus.Publish(ctx, content.PullEvents())
}
