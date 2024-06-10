package cr

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lorenzoMrt/ContentInsight/kit/event"
)

var ErrInvalidContentID = errors.New("invalid Content ID")

type ContentID struct {
	value string
}

// NewContentID instantiate the VO for ContentID
func NewContentID(value string) (ContentID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return ContentID{}, fmt.Errorf("%w: %s", ErrInvalidContentID, value)
	}

	return ContentID{
		value: v.String(),
	}, nil
}

func (id ContentID) String() string {
	return id.value
}

var ErrEmptyContentTitle = errors.New("The title can not be empty")

type ContentTitle struct {
	value string
}

// NewContentTitle instantiate VO for ContentTitle
func NewContentTitle(value string) (ContentTitle, error) {
	if value == "" {
		return ContentTitle{}, ErrEmptyContentTitle
	}

	return ContentTitle{
		value: value,
	}, nil
}

func (title ContentTitle) String() string {
	return title.value
}

type ContentRepository interface {
	Save(ctx context.Context, content Content) error
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=ContentRepository

type Content struct {
	id              ContentID
	title           ContentTitle
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

	events []event.Event
}
type Metadata struct {
	Views    int
	Likes    int
	Comments int
}

func NewContent(id string, title string, description string, contentType string, categories []string, tags []string, author string, publicationDate time.Time, contentUrl string, duration *int, language string, coverImage string, metadata Metadata, status string, source string, visibility string) (Content, error) {
	idVO, err := NewContentID(id)
	if err != nil {
		return Content{}, err
	}

	contentTitleVO, err := NewContentTitle(title)
	if err != nil {
		return Content{}, err
	}
	content := Content{
		id:              idVO,
		title:           contentTitleVO,
		description:     description,
		contentType:     contentType,
		categories:      categories,
		tags:            tags,
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
	content.Record(NewContentCreatedEvent(idVO.String(), contentTitleVO.String(), description, contentType, author, contentUrl, language, coverImage, status, source, visibility, categories, tags, metadata, duration, publicationDate))
	return content, nil
}

// Getters
func (c Content) ID() ContentID {
	return c.id
}
func (c Content) Title() ContentTitle {
	return c.title
}
func (c Content) Description() string {
	return c.description
}
func (c Content) ContentType() string {
	return c.contentType
}
func (c Content) Categories() []string {
	return c.categories
}
func (c Content) Tags() []string {
	return c.tags
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

func (c *Content) Record(evt event.Event) {
	c.events = append(c.events, evt)
}

func (c Content) PullEvents() []event.Event {
	evt := c.events
	c.events = []event.Event{}

	return evt
}
