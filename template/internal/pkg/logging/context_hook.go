package logging

import (
	"github.com/sirupsen/logrus"
)

type ContextHook struct {
	Keys []string
}

func NewContextHook(keys []string) ContextHook {
	return ContextHook{
		Keys: keys,
	}
}

func (hook ContextHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.DebugLevel,
		logrus.InfoLevel,
		logrus.WarnLevel,
		logrus.ErrorLevel,
		logrus.FatalLevel,
		logrus.PanicLevel,
	}
}

func (hook ContextHook) Fire(entry *logrus.Entry) error {
	if entry.Context == nil {
		return nil
	}

	for _, key := range hook.Keys {
		value := entry.Context.Value(key)
		if value != nil {
			entry.Data[key] = value
		}
	}
	return nil
}
