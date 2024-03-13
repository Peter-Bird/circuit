package main

import (
	"circuit/event"
	"circuit/flip"
	"circuit/sink"
	"circuit/source"
	"circuit/wire"
)

func main() {
	src := source.New("Power Source")
	w := wire.New("Wire1")
	sw := flip.New("Switch")
	snk := sink.New("Ground")

	w.Join(src)
	w.Join(sw)
	sw.Attach(snk)

	sw.Toggle()
	src.On()

	event.Trigger()
}
