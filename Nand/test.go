package Nand

import "fmt"

const (
	NandOut = "Nand: %s output is: %t\n"
)

func Test() {
	println("Testing...")

	nand := New("NAND01")
	fmt.Printf(NandOut, nand.GetLabel(), nand.Get()) // Expect true

	nand.Weld(hello)
	nand.Set(0, true)
	fmt.Printf(NandOut, nand.GetLabel(), nand.Get()) // Expect true

	nand.Set(1, true)
	fmt.Printf(NandOut, nand.GetLabel(), nand.Get()) // Expect false

	nand.Set(1, false)
}

func hello(val bool) {
	fmt.Println(val)
}
