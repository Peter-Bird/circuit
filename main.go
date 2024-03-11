package main

import (
	"fmt"

	"circuit/flip"
	"circuit/sink"
	"circuit/source"
	"circuit/wire"
)

func main() {
	src := source.New("Power Source")
	sw := flip.New("switch1")
	w := wire.New("Wire1")

	snk := sink.New("Ground")

	w.WeldTo(src)
	w.WeldTo(sw)
	sw.Attach(snk)

	src.On()

	// Initially, the switch is off, so the wire won't receive any signal from the source
	fmt.Println("Before flipping the switch:")
	fmt.Printf("Wire signal level: %v\n", w.Get()) // Get the initial state of the wire

	sw.Toggle() // Turn the switch on, allowing the signal to flow from the source through the wire

	fmt.Println("After flipping the switch on:")
	fmt.Printf("Wire signal level: %v\n", w.Get()) // Get the state of the wire after the switch is turned on

	// If you want to see the effect of turning the switch off again, you can toggle it and check the wire's state again
	sw.Toggle() // Turn the switch off

	fmt.Println("After flipping the switch off:")
	fmt.Printf("Wire signal level: %v\n", w.Get()) // Get the state of the wire after the switch is turned off
}
