package creating

import (
	"context"
	"errors"

	cr "github.com/lorenzoMrt/ContentInsight/internal"
	"github.com/lorenzoMrt/ContentInsight/internal/increasing"
	"github.com/lorenzoMrt/ContentInsight/kit/event"
)

type IncreaseContentsCounterOnContentCreated struct {
	increasingService increasing.ContentCounterService
}

func NewIncreaseContentsCounterOnContentCreated(increaserService increasing.ContentCounterService) IncreaseContentsCounterOnContentCreated {
	return IncreaseContentsCounterOnContentCreated{
		increasingService: increaserService,
	}
}

func (e IncreaseContentsCounterOnContentCreated) Handle(_ context.Context, evt event.Event) error {
	ContentCreatedEvt, ok := evt.(cr.ContentCreatedEvent)
	if !ok {
		return errors.New("unexpected event")
	}

	return e.increasingService.Increase(ContentCreatedEvt.ID())
}
