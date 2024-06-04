package cr

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type ContentRepository interface {
	Save(ctx context.Context, content Content) error
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=ContentRepository

type Content struct {
	uuid            uuid.UUID
	details         Details
	classification  Classification
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

type Details struct {
	title       string
	description string
	contentType string
}

type Classification struct {
	categories []string
	tags       []string
}

type Metadata struct {
	Views    int
	Likes    int
	Comments int
}

// Constructor
func NewContent(uuid uuid.UUID, details Details, classification Classification, author string, publicationDate time.Time, contentUrl string, duration *int, language string, coverImage string, metadata Metadata, status string, source string, visibility string) Content {
	return Content{
		uuid:            uuid,
		details:         details,
		classification:  classification,
		author:          author,
		publicationDate: publicationDate,
		contentURL:      contentUrl,
		duration:        duration,
		language:        language,
		coverImage:      coverImage,
		metadata:        metadata,
		status:          status,
		source:          source,
		visibility:      visibility,
	}
}

// Getters
func (c Content) UUID() uuid.UUID {
	return c.uuid
}

func (c Content) Details() Details {
	return c.details
}

func (c Content) Classification() Classification {
	return c.classification
}

func (c Content) Author() string {
	return c.author
}

func (c Content) PublicationDate() time.Time {
	return c.publicationDate
}

func (c Content) ContentURL() string {
	return c.contentURL
}

func (c Content) Duration() *int {
	return c.duration
}

func (c Content) Language() string {
	return c.language
}

func (c Content) CoverImage() string {
	return c.coverImage
}

func (c Content) Metadata() Metadata {
	return c.metadata
}

func (c Content) Status() string {
	return c.status
}

func (c Content) Source() string {
	return c.source
}

func (c Content) Visibility() string {
	return c.visibility
}
