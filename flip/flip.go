package flip

import (
	"circuit/core"
	"log"
)

type SwitchState bool

const (
	Off SwitchState = false
	On  SwitchState = true
)

type Switch struct {
	State SwitchState
	label string
	level core.SigType

	endPoints []core.Digital
}

func New(label string) *Switch {
	return &Switch{
		State: Off,
		level: core.Z,
		label: label,

		endPoints: []core.Digital{},
	}
}

func (s *Switch) Toggle() {
	s.State = !s.State
}

func (s *Switch) Attach(o core.Digital) {
	if len(s.endPoints) >= 2 {
		log.Println("Don't")
		return
	}
	s.endPoints = append(s.endPoints, o)
}

func (s *Switch) Set(level core.SigType, source core.Digital) {
	if s.State == On {
		if len(s.endPoints) == 2 {
			s.level = level

			for _, p := range s.endPoints {
				if p != source {
					p.Set(level, s)
				}
			}
		}
	}
}

func (s *Switch) Get() core.SigType { return s.level }
func (s *Switch) GetLabel() string  { return s.label }
