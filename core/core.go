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
	Weldable
	Flowable
}

// Weldable: objects that can be welded
// together.
type Weldable interface {
	WeldTo(other Joinable)
	Attach(other Joinable)
}

// Flowable: objects that can transmit
// flow.
type Flowable interface {
	Get() SigType
	Set(SigType)
}

// Joint is the result of welding
// conducting elements together
type Joint struct {
	Partners []Joinable
}
