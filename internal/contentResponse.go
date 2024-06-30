package cr

import (
	"time"

	"github.com/lorenzoMrt/ContentInsight/kit/event"
)

type ContentResponse struct {
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
	metadata        MetadataResponse
	status          string
	source          string
	visibility      string

	events []event.Event
}
type MetadataResponse struct {
	views    int
	likes    int
	comments int
}

func toMetadataResponse(metadata Metadata) MetadataResponse {
	return MetadataResponse{
		views:    metadata.Views,
		likes:    metadata.Likes,
		comments: metadata.Comments,
	}
}

func NewContentResponse(content Content) ContentResponse {
	return ContentResponse{
		id:              content.id.String(),
		title:           content.title.String(),
		description:     content.description,
		contentType:     content.contentType,
		categories:      content.categories,
		tags:            content.tags,
		author:          content.author,
		publicationDate: content.publicationDate,
		contentURL:      content.contentURL,
		duration:        content.duration,
		language:        content.language,
		coverImage:      content.coverImage,
		metadata:        toMetadataResponse(content.metadata),
		status:          content.status,
		source:          content.source,
		visibility:      content.visibility,
	}
}
