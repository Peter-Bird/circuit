package flip

import (
	"circuit/core"
)

type SwitchState bool

const (
	Off SwitchState = false
	On  SwitchState = true
)

type Switch struct {
	core.Joint
	label string
	level core.SigType
	State SwitchState
}

func New(label string) *Switch {
	return &Switch{
		Joint: core.Joint{},
		State: Off,
		level: core.Z,
		label: label,
	}
}

func (s *Switch) Toggle() {
	s.State = !s.State

	for _, p := range s.Partners {
		if s.State == On {
			p.Set(p.Get())
		} else {
			p.Set(core.Low)
		}
	}
}

func (s *Switch) Attach(o core.Joinable) {
	s.Partners = append(s.Partners, o)
	if s.State == On {
		o.Set(o.Get())
	} else {
		o.Set(core.Low)
	}
}

func (s *Switch) Set(z core.SigType) {}
func (s *Switch) Get() core.SigType  { return s.level }
func (s *Switch) GetLabel() string   { return s.label }
