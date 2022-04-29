package settings

import (
	"github.com/tsarbas/goenvparse"
)

type Settings struct {
}

func (s *Settings) Parse() error {
	return goenvparse.Parse(s, ".env")
}

func New() *Settings {
	s := &Settings{}
	return s
}
