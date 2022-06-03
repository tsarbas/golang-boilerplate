package bus

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
)

func logEvent(ctx context.Context, event Event) error {
	data, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("event marshal error: %w", err)
	}

	fields := logrus.Fields{}
	err = json.Unmarshal(data, &fields)
	if err != nil {
		return fmt.Errorf("event fields unmarshal error: %w", err)
	}

	logrus.WithContext(ctx).
		WithFields(fields).
		WithField("event", fmt.Sprintf("%T", event)).
		Info("Event received")
	return err
}
