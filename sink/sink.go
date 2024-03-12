package sink

import (
	"circuit/core"
)

var _ core.Digital = (*Sink)(nil)

// Sink represents a ground connection
type Sink struct {
	port  core.Digital
	label string
}

func New(label string) *Sink {
	return &Sink{
		label: label,
	}
}

func (s *Sink) Attach(w core.Digital)              {}
func (s *Sink) Set(l core.SigType, w core.Digital) {}

func (s *Sink) GetLabel() string { return s.label }
