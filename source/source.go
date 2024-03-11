package source

import (
	"circuit/core"
	"log"
)

type Source struct {
	core.Joint
	label string
}

func New(label string) *Source {
	return &Source{
		Joint: core.Joint{},
		label: label,
	}
}

func (s *Source) WeldTo(o core.Joinable) {
	s.Attach(o)
	o.Attach(s)
}

func (s *Source) Attach(o core.Joinable) {
	if len(s.Partners) >= 1 {
		log.Println("Don't")
		return
	}
	s.Partners = append(s.Partners, o)
	s.Set(core.High)
}

func (s *Source) Set(level core.SigType) {
	if len(s.Partners) == 1 {
		for _, p := range s.Partners {
			p.Set(core.High)
		}
	}
}

func (s *Source) Get() core.SigType { return core.High }
func (s *Source) GetLabel() string  { return s.label }

func (s *Source) On() { s.Set(core.High) }
