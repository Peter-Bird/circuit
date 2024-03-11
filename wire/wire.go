package wire

import (
	"log"

	"circuit/core"
)

/*
.             "label"
.	  ⚫━━━━━━━━━━━━━━━━━━━━⚫
'       level (High, Low, Z)
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
	if len(w.Partners) >= 2 {
		log.Println("Don't")
		return
	}
	w.Partners = append(w.Partners, o)
}

func (w *Wire) Set(level core.SigType) {
	if len(w.Partners) == 2 {
		old := w.level

		if level != old {
			w.level = level

			for _, p := range w.Partners {
				if p.Get() == old {
					p.Set(level)
				}
			}
		}
	}
}

func (w *Wire) Get() core.SigType { return w.level }
func (w *Wire) GetLabel() string  { return w.label }
