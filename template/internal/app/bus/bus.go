package bus

import (
	"context"
	"fmt"
)

type (
	Storage interface {
		Execute(context.Context, interface{}, interface{}) error
	}
	Event interface{}
)

func HandleEvent(
	ctx context.Context,
	event Event,
	store Storage,
) (interface{}, error) {
	err := logEvent(ctx, event)
	if err != nil {
		return event, fmt.Errorf("log event error: %w", err)
	}

	switch e := event.(type) {

	default:
		return e, nil
	}
}
