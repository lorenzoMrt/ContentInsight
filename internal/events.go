package cr

import (
	"time"

	"github.com/lorenzoMrt/ContentInsight/kit/event"
)

const ContentCreatedEventType event.Type = "events.Content.created"

type ContentCreatedEvent struct {
	event.BaseEvent
	id              string
	title           string
	description     string
	contentType     string
	categories      []string
	tags            []string
	author          string
	publicationDate time.Time
	contentURL      string
	duration        *int
	language        string
	coverImage      string
	metadata        Metadata
	status          string
	source          string
	visibility      string
}

func NewContentCreatedEvent(id, title, description, contentType, author, contentURL, language, coverImage, status, source, visibility string, categories, tags []string, metadata Metadata, duration *int, publicationDate time.Time) ContentCreatedEvent {
	return ContentCreatedEvent{
		id:              id,
		title:           title,
		description:     description,
		contentType:     contentType,
		categories:      categories,
		tags:            tags,
		author:          author,
		publicationDate: publicationDate,
		contentURL:      contentURL,
		duration:        duration,
		language:        language,
		coverImage:      coverImage,
		metadata:        metadata,
		status:          status,
		source:          source,
		visibility:      visibility,

		BaseEvent: event.NewBaseEvent(id),
	}
}

func (e ContentCreatedEvent) Type() event.Type {
	return ContentCreatedEventType
}

func (e ContentCreatedEvent) ContentID() string {
	return e.id
}

func (e ContentCreatedEvent) ContentTitle() string {
	return e.title
}

func (e ContentCreatedEvent) ContentDescription() string {
	return e.description
}
func (e ContentCreatedEvent) ContentType() string {
	return e.contentType
}
func (e ContentCreatedEvent) ContentCategories() []string {
	return e.categories
}
func (e ContentCreatedEvent) ContentTags() []string {
	return e.tags
}
func (e ContentCreatedEvent) ContentAuthor() string {
	return e.author
}
func (e ContentCreatedEvent) ContentPublicationDate() time.Time {
	return e.publicationDate
}
func (e ContentCreatedEvent) ContentURL() string {
	return e.contentURL
}
func (e ContentCreatedEvent) ContentDuration() *int {
	return e.duration
}
func (e ContentCreatedEvent) ContentLanguage() string {
	return e.language
}
func (e ContentCreatedEvent) ContentCoverImage() string {
	return e.coverImage
}
func (e ContentCreatedEvent) ContentMetadata() Metadata {
	return e.metadata
}
func (e ContentCreatedEvent) ContentStatus() string {
	return e.status
}
func (e ContentCreatedEvent) ContentSource() string {
	return e.source
}

func (e ContentCreatedEvent) ContentVisibility() string {
	return e.visibility
}
