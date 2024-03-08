package nand

import (
	"fmt"

	"circuit/port"
)

/*
.     ( ⊼ )
.   ┏━━━━━━━━┓
.  ●┫A      Y┣●
.  ●┫B       ┃
.   ┃… NAND  ┃
.  ●┫E       ┃
.  ●┫F       ┃
.   ┗━━━━━━━━┛
*/

const (
	BadGate = "NAND gate requires at least 2 inputs"
	BadPort = "Port idx out of range"
)

type Nand struct {
	ports []port.Pin
}

func New(inputs int) *Nand {

	if inputs < 2 {
		panic(BadGate)
	}

	ports := []port.Pin{}

	for i := 0; i < inputs; i++ {
		letter := fmt.Sprintf("%c", 'A'+i)
		ports = append(ports, port.New(letter))
	}

	ports = append(ports, port.New("Y"))

	return &Nand{ports: ports}
}

func (n *Nand) Set(idx int, level bool) {

	if idx < 0 || idx >= len(n.ports)-1 {
		panic(BadPort)
	}

	n.ports[idx].Set(level)
}

func (n *Nand) Logic() {

	n.ports[len(n.ports)-1].Set(true)

	for i := 0; i < len(n.ports)-1; i++ {
		if !n.ports[i].Get() {
			return
		}
	}

	n.ports[len(n.ports)-1].Set(false)
}

func (n *Nand) Get() bool {
	return n.ports[len(n.ports)-1].Get()
}

func (n *Nand) GetLabel(i int) string {
	return n.ports[i].GetLabel()
}
