package wire_test

import (
	"fmt"
	"testing"

	"circuit/core"
	"circuit/wire"
)

func TestT1(t *testing.T) {
	const imgT1 = `
         Wire1
    ⚫━━━━━━━━━━━━━⚫
`
	fmt.Print(imgT1)

	fmt.Print("\nMake:\t\t")
	wire1 := wire.New("Wire1")
	Display(wire1)

	fmt.Print("Set High:\t")
	wire1.Set(core.High)
	Display(wire1)

	fmt.Print("Set Low:\t")
	wire1.Set(core.Low)
	Display(wire1)

	fmt.Println()
}

func TestT2(t *testing.T) {
	const imgT2 = `
    E[0]     E[1]     E[0]     E[1]
    ↓  Wire1 ↓  Wire2 ↓  Wire3 ↓
    ⚫━━━━━━━⚪━━━━━━━⚪━━━━━━━⚫
              ↑        ↑
            E[0]       E[1]
`
	fmt.Print(imgT2)

	fmt.Print("\nMake:\t\t")
	wire1 := wire.New("Wire1")
	Display(wire1)

	fmt.Print("Make:\t\t")
	wire2 := wire.New("Wire2")
	Display(wire2)

	fmt.Print("Make:\t\t")
	wire3 := wire.New("Wire3")
	Display(wire3)

	fmt.Println("\nWeld:")
	wire1.WeldTo(wire2)
	wire2.WeldTo(wire3)
	Display(wire1)
	Display(wire2)
	Display(wire3)

	fmt.Println("\nSet Wire1 High:")
	wire1.Set(core.High)
	Display(wire1)
	Display(wire2)
	Display(wire3)

	fmt.Println("\nSet Wire2 Low:")
	wire2.Set(core.Low)
	Display(wire1)
	Display(wire2)
	Display(wire3)
}

func TestT3(t *testing.T) {
	const imgT2 = `
      Wire1    Wire2      
   ⚫━━━━━━━⚪━━━━━━━⚪
                         ┃ Wire3
      Wire5    Wire4     ┃
   ⚫━━━━━━━⚪━━━━━━━⚪
`
	fmt.Print(imgT2)

	fmt.Print("\nMake:\t\t")
	wire1 := wire.New("Wire1")
	wire2 := wire.New("Wire2")
	wire3 := wire.New("Wire3")
	wire4 := wire.New("Wire4")
	wire5 := wire.New("Wire5")
	wire1.WeldTo(wire2)
	wire2.WeldTo(wire3)
	wire3.WeldTo(wire4)
	wire4.WeldTo(wire5)

	// Open circuit
	fmt.Println("\nSet Wire1 High:")
	wire1.Set(core.High)
	Display(wire1)
	Display(wire2)
	Display(wire3)
	Display(wire4)
	Display(wire5)

	// Propogate
	fmt.Println("\n\nSet Wire2 High:")
	wire2.Set(core.High)
	Display(wire1)
	Display(wire2)
	Display(wire3)
	Display(wire4)
	Display(wire5)

	// Back Propogate
	fmt.Println("\n\nSet Wire3 Low:")
	wire3.Set(core.Low)
	Display(wire1)
	Display(wire2)
	Display(wire3)
	Display(wire4)
	Display(wire5)
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

func Display[T core.Stringer](w T) {
	fmt.Println(w.String())
}
