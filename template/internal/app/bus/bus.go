package bus

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type (
	Bus struct {
		Validator *validator.Validate
	}
	Event   interface{}
	Command interface{}
)

func NewBus(s Storage) *Bus {
	return &Bus{
		Validator: validator.New(),
	}
}

func (b *Bus) HandleEvents(events []Event) {
	for _, event := range events {
		_, err := b.HandleEvent(event)
		if err != nil {
			logrus.Error(fmt.Errorf("Handle event error: %w", err).Error())
		}
	}
}

func (b *Bus) HandleEvent(event Event) (interface{}, error) {
	switch e := event.(type) {

	default:
		return e, errors.New("Invalid event")
	}
}

func (b *Bus) HandleCommand(command Command) (interface{}, error) {

	err := b.Validator.Struct(command)
	if err != nil {
		return nil, err
	}

	switch c := command.(type) {

	default:
		return с, errors.New("Invalid command")
	}
}
