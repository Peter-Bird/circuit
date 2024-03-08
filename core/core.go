package core

// SigType is used to represent
// the state of a digital signal.
type SigType int

const (
	Low  SigType = iota // 0
	High                // 1
	Z                   // High-Impedance
)

// Weldable: objects that can be welded
// together.
type Weldable interface {
	WeldTo(other Weldable)
	Attach(other Weldable)
}

type Joint struct {
	Partners []Weldable
}

func (j *Joint) WeldTo(o Weldable) {
	j.Attach(o)
	o.Attach(j)
}

func (j *Joint) Attach(o Weldable) {
	j.Partners = append(j.Partners, o)
}
