package creating

import (
	"context"
	"errors"
	"time"

	cr "github.com/lorenzoMrt/ContentInsight/internal"
	"github.com/lorenzoMrt/ContentInsight/kit/command"
)

const ContentCommandType command.Type = "command.creating.Content"

// ContentCommand is the command dispatched to create a new Content.
type ContentCommand struct {
	id              string
	title           string
	description     string
	contentType     string
	categories      []string
	tags            []string
	author          string
	publicationDate time.Time
	contentUrl      string
	duration        *int
	language        string
	coverImage      string
	metadata        Metadata
	status          string
	source          string
	visibility      string
}

type Metadata struct {
	Views    int
	Likes    int
	Comments int
}

func toCrMetadata(metadata Metadata) cr.Metadata {
	return cr.Metadata{
		Views:    metadata.Views,
		Likes:    metadata.Likes,
		Comments: metadata.Comments,
	}
}

// NewContentCommand creates a new ContentCommand.
func NewContentCommand(id string, title string, description string, contentType string, categories []string, tags []string, author string, publicationDate time.Time, contentUrl string, duration *int, language string, coverImage string, metadata Metadata, status string, source string, visibility string) ContentCommand {
	return ContentCommand{
		id:              id,
		title:           title,
		description:     description,
		contentType:     contentType,
		categories:      categories,
		tags:            tags,
		author:          author,
		publicationDate: publicationDate,
		contentUrl:      contentUrl,
		duration:        duration,
		language:        language,
		coverImage:      coverImage,
		metadata:        metadata,
		status:          status,
		source:          source,
		visibility:      visibility,
	}
}

func (c ContentCommand) Type() command.Type {
	return ContentCommandType
}

// ContentCommandHandler is the command handler
// responsible for creating Contents.
type ContentCommandHandler struct {
	service ContentService
}

// NewContentCommandHandler initializes a new ContentCommandHandler.
func NewContentCommandHandler(service ContentService) ContentCommandHandler {
	return ContentCommandHandler{
		service: service,
	}
}

// Handle implements the command.Handler interface.
func (h ContentCommandHandler) Handle(ctx context.Context, cmd command.Command) error {
	createContentCmd, ok := cmd.(ContentCommand)
	if !ok {
		return errors.New("unexpected command")
	}

	return h.service.CreateContent(
		ctx,
		createContentCmd.id,
		createContentCmd.title,
		createContentCmd.description,
		createContentCmd.contentType,
		createContentCmd.categories,
		createContentCmd.tags,
		createContentCmd.author,
		createContentCmd.publicationDate,
		createContentCmd.contentUrl,
		createContentCmd.duration,
		createContentCmd.language,
		createContentCmd.coverImage,
		toCrMetadata(createContentCmd.metadata),
		createContentCmd.status,
		createContentCmd.source,
		createContentCmd.visibility,
	)
}
