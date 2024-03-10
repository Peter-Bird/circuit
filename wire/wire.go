package wire

import (
	"fmt"
	"strings"

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
		Joint: core.Joint{},
		label: label,
		level: core.Z,
	}
}

func (w *Wire) WeldTo(o core.Joinable) {
	w.Attach(o)
	o.Attach(w)
}

func (w *Wire) Attach(o core.Joinable) {
	w.Partners = append(w.Partners, o)
}

func (w *Wire) Set(level core.SigType) {
	if len(w.Partners) == 2 {
		old := w.level

		if level != old {
			w.level = level

			for _, partner := range w.Partners {
				if partner.Get() == old {
					partner.Set(level)
				}
			}
		}
	}
}

func (w *Wire) Get() core.SigType { return w.level }
func (w *Wire) GetLabel() string  { return w.label }

func (w *Wire) String() string {
	var builder strings.Builder

	fmt.Fprintf(&builder, "%s is %v, ", w.label, w.level)

	for idx, part := range w.Partners {
		// Attempt to assert the type of part to *Wire
		if wirePart, ok := part.(*Wire); ok {
			// If successful, wirePart is now a *Wire and we can access its label
			fmt.Fprintf(&builder, "P: %d, is: %v, ", idx, wirePart.label)
		} else {
			// If not successful, part is not a *Wire and you can handle accordingly
			fmt.Fprintf(&builder, "P: %d, is of a different type\n", idx)
		}
	}

	return builder.String()
}
