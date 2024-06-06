package creating

import (
	"context"
	"time"

	cr "github.com/lorenzoMrt/ContentInsight/internal"
)

type ContentService struct {
	contentRepository cr.ContentRepository
}

func NewContentService(contentRepository cr.ContentRepository) ContentService {
	return ContentService{
		contentRepository: contentRepository,
	}
}

func (s ContentService) CreateContent(ctx context.Context, uuid string, title string, description string, contentType string, categories []string, tags []string, author string, publicationDate time.Time, contentURL string, duration *int, language string, coverImage string, metadata cr.Metadata, status string, source string, visibility string) error {
	content, err := cr.NewContent(uuid, title, description, contentType, categories, tags, author, publicationDate, contentURL, duration, language, coverImage, metadata, status, source, visibility)

	if err != nil {
		return err
	}

	return s.contentRepository.Save(ctx, content)
}
