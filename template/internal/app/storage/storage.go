package storage

import (
	"context"
	"errors"
)

type Storage struct{}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) Execute(
	ctx context.Context,
	query interface{},
	dest interface{},
) error {
	switch q := query.(type) {

	default:
		return errors.New("invalid query")
	}
}
