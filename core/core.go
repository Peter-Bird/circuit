package core

// SigType is used to represent
// the state of a digital signal.
type SigType int

const (
	Low  SigType = iota // 0
	High                // 1
	Z                   // High-Impedance
)

type Joinable interface {
	WeldTo(other Joinable)
	Attach(other Joinable)

	Get() SigType
	Set(SigType)

	GetLabel() string
	Stringer
}

// Joint is the result of welding
// conducting elements together
type Joint struct {
	Partners []Joinable
}

type Stringer interface {
	String() string
}
