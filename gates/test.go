package gates

import (
	"fmt"

	"circuit/gates/nand"
)

const (
	Output = "Output state of NAND gate: %t.\n"
)

func TestNand() {

	ports := 2

	gate := nand.New(ports)

	gate.Logic()
	fmt.Printf(Output, gate.Get())

	gate.Set(0, true)
	gate.Logic()
	fmt.Printf(Output, gate.Get())

	gate.Set(1, true)
	gate.Logic()
	fmt.Printf(Output, gate.Get())

	for i := 0; i < 3; i++ {
		fmt.Printf("%v\n", gate.GetLabel(i))
	}
}
