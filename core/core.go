package core

import "fmt"

// SigType is used to represent
// the state of a digital signal.
type SigType int

const (
	Low  SigType = iota // 0
	High                // 1
	Z                   // High-Impedance
)

type Connection interface {
	Weldable
	Conducter
}

// Weldable: objects that can be welded
// together.
type Weldable interface {
	WeldTo(other Connection)
	Attach(other Connection)
}

type Joint struct {
	Partners []Connection
}

func (j *Joint) WeldTo(o Connection) {
	fmt.Printf("WeldTo %v  %v\n", j, o)
	j.Attach(o)
	o.Attach(j)
	fmt.Printf("Attached %v  %v\n", j, o)
}

func (j *Joint) Attach(o Connection) {
	fmt.Printf("Attaching %v\n", j)
	j.Partners = append(j.Partners, o)
}

type Conducter interface {
	Get() SigType
	Set(SigType)
}

func (j *Joint) Get() SigType {

	fmt.Print("In Get\n")
	if len(j.Partners) > 0 {
		return j.Partners[0].Get()
	}

	return Z
}

func (j *Joint) Set(sig SigType) {
	for _, partner := range j.Partners {
		partner.Set(sig)
	}
}
