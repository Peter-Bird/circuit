package main

import (
	"circuit/sink"
	"circuit/source"
	"circuit/wire"
)

func main() {
	src := source.New("Power Source")
	w := wire.New("Wire1")

	snk := sink.New("Ground")

	w.Join(src)
	w.Join(snk)

	src.On()
}
