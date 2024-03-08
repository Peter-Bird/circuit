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

func (w *Wire) Weld(other core.Connection) {
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

		for _, partner := range w.Partners {
			fmt.Printf("Partner %v\n", partner)
			if partner.Get() == old {
				fmt.Print("Setting Partner\n")
				partner.Set(level)
			}
		}
	}
}

func (w *Wire) Get() core.SigType { return w.level }
func (w *Wire) GetLabel() string  { return w.label }

func (w *Wire) String() {
	fmt.Printf("%s is %d, ", w.label, w.level)

	for idx, part := range w.Partners {
		fmt.Printf("P: %d, is: %v\n", idx, part)
	}

	fmt.Println()
}
