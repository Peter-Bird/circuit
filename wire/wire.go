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

var _ core.Digital = (*Wire)(nil)

type Wire struct {
	label string
	level core.SigType

	endPoints []core.Digital
}

func New(label string) *Wire {
	return &Wire{
		label: label,
		level: core.Z,

		endPoints: []core.Digital{},
	}
}

func (w *Wire) Join(o core.Digital) {
	w.Attach(o)
	o.Attach(w)
}

func (w *Wire) Attach(o core.Digital) {
	if len(w.endPoints) >= 2 {
		log.Println("Don't")
		return
	}
	w.endPoints = append(w.endPoints, o)
}

func (w *Wire) Set(level core.SigType, source core.Digital) {
	if len(w.endPoints) == 2 {
		w.level = level

		for _, p := range w.endPoints {
			if p != source {
				p.Set(level, w)
			}
		}
	}
}

func (w *Wire) Get() core.SigType { return w.level }
