package sink

import (
	"circuit/core"
	"log"
)

// Sink represents a ground connection
// or signal-consuming component
type Sink struct {
	core.Joint
	label string
}

func New(label string) *Sink {
	return &Sink{
		Joint: core.Joint{},
		label: label,
	}
}

func (s *Sink) Attach(o core.Joinable) {
	if len(s.Partners) >= 1 {
		log.Println("Don't")
		return
	}
	s.Partners = append(s.Partners, o)
}

func (s *Sink) Set(z core.SigType) {}
func (s *Sink) Get() core.SigType  { return core.Low }

func (s *Sink) GetLabel() string { return s.label }
