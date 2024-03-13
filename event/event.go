package event

import (
	"container/list"
	"fmt"
)

// Define the Element interface
type Element interface {
	ProcessEvent()
}

// Define a simple resistor as an example element
type Resistor struct {
	next []Element // Elements that follow this resistor in the circuit
}

func (r *Resistor) ProcessEvent() {
	fmt.Println("Resistor processing event")
	// Trigger events in the next elements
	for _, elem := range r.next {
		eventQueue.PushBack(elem)
	}
}

var eventQueue = list.New() // Global event queue

func Trigger() {
	// Example setup: resistor1 -> resistor2 -> resistor3
	resistor1 := &Resistor{}
	resistor2 := &Resistor{}
	resistor3 := &Resistor{}

	// Setting up the "circuit"
	resistor1.next = append(resistor1.next, resistor2)
	resistor2.next = append(resistor2.next, resistor3)

	// Start the simulation by adding an initial event
	eventQueue.PushBack(resistor1)

	// Event loop
	for eventQueue.Len() > 0 {
		// Get the next event from the queue
		front := eventQueue.Front()
		element := front.Value.(Element)

		// Process the event
		element.ProcessEvent()

		// Remove the processed event from the queue
		eventQueue.Remove(front)
	}
}
