package main

import (
	"fmt"
	// "circuit/Nand"
	// "circuit/gates"
	//"circuit/wire"
)

func main() {
	fmt.Println()
	title("WIRE TESTS")
	//wire.Test()

	// title("Gate NAND")
	// gates.TestNand()

	// title("Chip NAND")
	// Nand.Test()
}

func title(name string) {
	fmt.Println("\t====\t", name, "\t====")
}
