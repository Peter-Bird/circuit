package source

import (
	"circuit/core"
)

var _ core.Digital = (*Source)(nil)

type Source struct {
	port  core.Digital
	label string
}

func New(label string) *Source {
	return &Source{
		label: label,
	}
}

func (s *Source) On()  { s.Set(core.High, s) }
func (s *Source) Off() { s.Set(core.Z, s) }

func (s *Source) Attach(wire core.Digital) { s.port = wire }

func (s *Source) Set(level core.SigType, source core.Digital) {
	s.port.Set(core.High, s)
}
