package wire_test

import (
	"fmt"
	"testing"

	"circuit/core"
	"circuit/wire"
)

// TestNewWire checks if a new wire
// is correctly initialized.
func TestNewWire(t *testing.T) {
	fmt.Print(`
     Wire 1    
⚫━━━━━━━━━━━━⚫

`)
	name := "TestWire"
	w := wire.New(name)

	label := w.GetLabel()
	if label != name {
		t.Errorf(NameErr, name, label)
	}

	level := w.Get()
	if level != core.Z {
		t.Errorf(LevelErr, level)
	}
}

// TestSetNewWire assures that
// a disconnected wire doen't conduct.
func TestSetNewWire(t *testing.T) {
	fmt.Print(`
     Wire 1    
⚫━━━━━━━━━━━━⚫

`)
	w := wire.New("Wire 1")
	w.Set(core.High)

	level := w.Get()
	if level != core.Z {
		t.Errorf(LevelErr, level)
	}
}

func TestWeld(t *testing.T) {
	fmt.Print(`
   Wire 1    Wire 2
⚫━━━━━━━━⚪━━━━━━━━⚫

`)
	w1 := wire.New("Wire 1")
	w2 := wire.New("Wire 2")
	w1.WeldTo(w2)

	len1 := len(w1.Partners)
	if len1 != 1 {
		t.Errorf(WeldErr)
	}

	len2 := len(w2.Partners)
	if len2 != 1 {
		t.Errorf(WeldErr)
	}

	if w1.Partners[0] != w2 {
		t.Errorf(WrongPartners)
	}

	if w2.Partners[0] != w1 {
		t.Errorf(WrongPartners)
	}
}

// func TestT2(t *testing.T) {
// 	fmt.Print(`
// ↓  Wire 1↓  Wire 2↓  Wire 3↓
// ⚫━━━━━━━⚪━━━━━━━⚪━━━━━━━⚫

// `)

// 	wire1 := wire.New("Wire 1")
// 	wire2 := wire.New("Wire 2")
// 	wire3 := wire.New("Wire 3")

// 	fmt.Println("\nWeld:")
// 	wire1.WeldTo(wire2)
// 	wire2.WeldTo(wire3)

// 	fmt.Println("\nSet Wire1 High:")
// 	wire1.Set(core.High)

// 	fmt.Println("\nSet Wire2 Low:")
// 	wire2.Set(core.Low)

// }

// func TestT3(t *testing.T) {
// 	const imgT2 = `
//       Wire 1   Wire 2
//    ⚫━━━━━━━⚪━━━━━━━⚪
//                          ┃ Wire
//       Wire 5   Wire 4    ┃  3
//    ⚫━━━━━━━⚪━━━━━━━⚪
// `
// 	fmt.Print(imgT2)

// 	fmt.Print("\nMake:\t\t")
// 	wire1 := wire.New("Wire1")
// 	wire2 := wire.New("Wire2")
// 	wire3 := wire.New("Wire3")
// 	wire4 := wire.New("Wire4")
// 	wire5 := wire.New("Wire5")
// 	wire1.WeldTo(wire2)
// 	wire2.WeldTo(wire3)
// 	wire3.WeldTo(wire4)
// 	wire4.WeldTo(wire5)

// 	// Open circuit
// 	fmt.Println("\nSet Wire1 High:")
// 	wire1.Set(core.High)

// 	// Propogate
// 	fmt.Println("\n\nSet Wire2 High:")
// 	wire2.Set(core.High)

// 	// Back Propogate
// 	fmt.Println("\n\nSet Wire3 Low:")
// 	wire3.Set(core.Low)

// }

// func TestT4(t *testing.T) {
// 	const imgT2 = `
//          Wire 2
//       ⚪━━━━━━━⚪
// Wire  ┃           ┃ Wire
//   1   ┃           ┃  3
//       ⚪━━━━━━━⚪
//          Wire 4
// `
// 	fmt.Print(imgT2)

// 	fmt.Print("\nMake:\t\t")
// 	wire1 := wire.New("Wire1")
// 	wire2 := wire.New("Wire2")
// 	wire3 := wire.New("Wire3")
// 	wire4 := wire.New("Wire4")

// 	wire1.WeldTo(wire2)
// 	wire2.WeldTo(wire3)
// 	wire3.WeldTo(wire4)
// 	wire4.WeldTo(wire1)

// 	// Closed circuit (Loop)
// 	fmt.Println("\nSet Wire1 High:")
// 	wire1.Set(core.High)

// }
