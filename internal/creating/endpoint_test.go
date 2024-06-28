package creating

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/lorenzoMrt/ContentInsight/kit/command/commandmocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMakeCreateContentEndpoint(t *testing.T) {
	mockBus := new(commandmocks.Bus)
	mockBus.On(
		"Dispatch",
		mock.Anything,
		mock.AnythingOfType("creating.ContentCommand"),
	).Return(nil)

	endpoint := makeCreateContentEndpoint(mockBus)

	t.Run("Successful content creation", func(t *testing.T) {
		req := contentRequest{
			Uuid:            uuid.New().String(),
			Title:           "Test Title",
			Description:     "Test Description",
			ContentType:     "article",
			Categories:      []string{"Category1", "Category2"},
			Tags:            []string{"Tag1", "Tag2"},
			Author:          "Test Author",
			PublicationDate: time.Now(),
			ContentURL:      "https://example.com",
			Duration:        nil,
			Language:        "en",
			CoverImage:      "https://example.com/image.jpg",
			Metadata: metadataRequest{
				Views:    100,
				Likes:    10,
				Comments: 5,
			},
			Status:     "published",
			Source:     "Test Source",
			Visibility: "public",
		}

		resp, err := endpoint(context.Background(), req)
		assert.NoError(t, err)
		response := resp.(createContentResponse)
		assert.Nil(t, response.Err)
		mockBus.AssertCalled(t, "Dispatch", mock.Anything, mock.AnythingOfType("creating.ContentCommand"))
	})

}
