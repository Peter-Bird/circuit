package core

// SigType is used to represent
// the state of a digital signal.
type SigType int

const (
	Low SigType = iota
	High
	Z // High-Impedance
)

type Joinable interface {
	//	WeldTo(other Joinable)
	Attach(other Joinable)

	Get() SigType
	Set(SigType)

	GetLabel() string
}

// Joint is the result of welding
// conducting elements together
type Joint struct {
	Partners []Joinable
}
