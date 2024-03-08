package Nand

import (
	"circuit/gates/nand"
	"circuit/port"
)

/*
.             Nand (⊼)
.       ┏━━━━━━━━━━━━━━━━━━┓
.       ┃     ┏━━━━━━┓     ┃
.  A1 ━●┫━0 ━▶┫      ┣━ 2━▶┣●━ Y1
.       ┃     ┃ NAND ┃     ┃
.  B1 ━●┫━1 ━▶┫      ┃     ┃
.       ┃     ┗━━━━━━┛     ┃
.	    ┗━━━━━━━━━━━━━━━━━━┛
*/

type Nand struct {
	Label string

	Ports  []port.Port
	Module nand.Nand

	Change func(bool)
}

func New(label string) *Nand {

	chip := &Nand{
		Label: label,
		Ports: []port.Port{
			*(port.New("A")),
			*(port.New("B")),
			*(port.New("Y")),
		},
		Module: *(nand.New(2)),
		Change: nil,
	}
	chip.logic()

	return chip
}

func (l *Nand) Set(pNum int, level bool) {
	l.Module.Set(pNum, level)
	l.Ports[pNum].Set(level)

	old := l.Ports[2].Get()
	l.logic()

	new := l.Ports[2].Get()
	if old != new {
		if l.Change != nil {
			l.Change(new)
		}
	}
}

func (l *Nand) logic() {
	l.Module.Logic()
	l.Ports[2].Set(l.Module.Get())
}

func (l *Nand) Get() bool {
	return l.Ports[2].Get()
}

func (l *Nand) GetLabel() string {
	return l.Label
}

func (l *Nand) Weld(alert func(bool)) {
	l.Change = alert
}
