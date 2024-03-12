package core

// SigType is used to represent
// the state of a digital signal.
type SigType int

const (
	Low SigType = iota
	High
	Z // High-Impedance
)

type Digital interface {
	Attach(Digital)
	Set(SigType, Digital)
}

// // Port is the connection
// // point to the Source.
// type Port struct {
// 	Label    string
// 	EndPoint Digital
// }
