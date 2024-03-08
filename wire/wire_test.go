package wire_test

import (
	"fmt"
	"testing"

	"circuit/core"
	"circuit/wire"
)

/*
   .    W1-P1          W1-P2
   .     ↓     Wire1    ↓     Wire2
   .	  ⚪━━━━━━━━━━━━━⚪━━━━━━━━━━━━━⚪
   .                     ↑              ↑
   .                    W2-P1          W2-P2
*/

func TestT1(t *testing.T) {
	const imgT1 = `
         Wire1
    ⚫━━━━━━━━━━━━━⚫
`
	fmt.Print(imgT1)

	fmt.Print("\nMake:\t\t")
	wire1 := wire.New("Wire1")
	wire1.String()

	fmt.Print("Set High:\t")
	wire1.Set(core.High)
	wire1.String()

	fmt.Print("Set Low:\t")
	wire1.Set(core.Low)
	wire1.String()

	fmt.Println()
}

func TestT2(t *testing.T) {
	const imgT2 = `
    E[0]          E[1]
    ↓     Wire1    ↓     Wire2
    ⚫━━━━━━━━━━━━━⚪━━━━━━━━━━━━━⚫
                    ↑              ↑
                   E[0]          E[1]
`
	fmt.Print(imgT2)

	fmt.Print("\nMake:\t\t")
	wire1 := wire.New("Wire1")
	wire1.String()

	fmt.Print("Make:\t\t")
	wire2 := wire.New("Wire2")
	wire2.String()

	fmt.Println("\nWeld:")
	wire1.Weld(wire2)
	wire1.String()
	wire2.String()

	fmt.Println("\nSet Wire1 High:")
	wire1.Set(core.High)
	wire1.String()
	wire2.String()

	fmt.Println("\nSet Wire2 Low:")
	wire2.Set(core.Low)
	wire1.String()
	wire2.String()
}

// TestNewWire checks if a new wire is correctly initialized.
func TestNewWire(t *testing.T) {
	w := wire.New("TestWire")
	if w.GetLabel() != "TestWire" {
		t.Errorf("Expected wire label to be 'TestWire', got '%s'", w.GetLabel())
	}
	if w.Get() != core.Low {
		t.Errorf("Expected new wire level to be Low, got %d", w.Get())
	}
}
