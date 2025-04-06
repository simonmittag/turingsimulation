package main

import (
	"fmt"
)

// Symbol is a character on the tape

// TuringMachine represents the state of the machine

func main() {
	// Tape layout: A, B, _, _, SUM, CARRY
	tape := []Symbol{One, One, Blank, Blank} // A = 1, B = 1

	transitions := map[Key]Transition{
		// Read A
		{"q0", Zero}: {Zero, Right, "qStoreA0"},
		{"q0", One}:  {One, Right, "qStoreA1"},

		// A = 0
		{"qStoreA0", Zero}: {Zero, Right, "qXOR0"},
		{"qStoreA0", One}:  {One, Right, "qXOR1"},

		// A = 1
		{"qStoreA1", Zero}: {Zero, Right, "qXOR2"},
		{"qStoreA1", One}:  {One, Right, "qXOR3"},

		// XOR0: A=0, B=0 -> SUM=0, CARRY=0
		{"qXOR0", Blank}: {Zero, Right, "qAND0"},
		{"qAND0", Blank}: {Zero, Right, "HALT"},

		// XOR1: A=0, B=1 -> SUM=1, CARRY=0
		{"qXOR1", Blank}: {One, Right, "qAND1"},
		{"qAND1", Blank}: {Zero, Right, "HALT"},

		// XOR2: A=1, B=0 -> SUM=1, CARRY=0
		{"qXOR2", Blank}: {One, Right, "qAND2"},
		{"qAND2", Blank}: {Zero, Right, "HALT"},

		// XOR3: A=1, B=1 -> SUM=0, CARRY=1
		{"qXOR3", Blank}: {Zero, Right, "qAND3"},
		{"qAND3", Blank}: {One, Right, "HALT"},
	}

	tm := &TuringMachine{
		Tape:    tape,
		Head:    0,
		State:   "q0",
		Program: transitions,
	}
	tm.Run()

	fmt.Printf("Final Tape: %v\n", string(tm.Tape))
}
