package port

import (
	"fmt"
)

type Pin interface {
	Set(bool)
	Get() bool
	GetLabel() string
	Weld(int, Pin)
}

/*
.   Port:
.   -----
.
.	 ⚪     (⚫) ← Level
.    ↑
.   Label
*/

const (
	Passport = "%s is %t\n"
)

// Port is a simple implementation
// of the Pin interface.
type Port struct {
	label string
	level bool

	joints []Pin
}

// Creates a new Port
func New(label string) *Port {
	return &Port{label: label}
}

// Set changes the level of the Port.
func (p *Port) Set(level bool) {
	p.level = level
}

// Get reads the level of the Port.
func (p *Port) Get() bool {
	return p.level
}

// GetLabel reads the label of the Port.
func (p *Port) GetLabel() string {
	return p.label
}

func (p *Port) Weld(edge int, pin Pin) {
	p.joints = append(p.joints, pin)
}

func (p *Port) Print() {
	fmt.Printf(Passport, p.label, p.level)
}
