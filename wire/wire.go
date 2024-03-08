package wire

import (
	"fmt"

	"circuit/core"
)

/*
.           label
.	  ⚫━━━━━━━━━━━━━⚫
.     ↑              ↑
.   edge[0]        edge[1]
*/

type Wire struct {
	core.Joint
	label string
	level core.SigType
}

func New(label string) *Wire {
	return &Wire{
		label: label,
		level: core.Z,
	}
}

func (w *Wire) Weld(other core.Weldable) {
	w.WeldTo(other)

	// An attached Wire can transmit current
	if len(w.Partners) == 2 {
		w.level = core.Low // Todo: Clarify
	}
}

func (w *Wire) Set(level core.SigType) {
	old := w.level

	if level != old {
		w.level = level

		// for i := 0; i < 2; i++ {
		// 	edge := w.edges[i]
		// 	if edge != nil {
		// 		if edge.Get() == old {
		// 			edge.Set(level)
		// 		}
		// 	}
		// }
	}
}

func (w *Wire) Get() core.SigType { return w.level }
func (w *Wire) GetLabel() string  { return w.label }

func (w *Wire) String() {
	fmt.Printf("%v\n", w)
}
