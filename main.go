package main

import (
	"fmt"
)

// main is the entry point of the application that initializes the Turing machine,
// loads its program, runs it, and outputs the final tape state.
func main() {
	//we initialise the tape
	tape := initialiseTape() // A = 1, B = 1

	//we load our program
	program := getProgram()

	//the machine is initialised with a tape, a starter position for the head, an initial state, and the actual program
	tm := &TuringMachine{
		Tape:    tape,
		Head:    0,
		Program: program,
		State:   "q0",
	}

	//run the machine
	tm.Run()

	//after the machine stops, print the contents of the tape
	fmt.Printf("Final Tape: %v\n", string(tm.Tape))
}

// initialiseTape initializes the Turing machine tape with predefined symbols: two '1's followed by two blanks.
func initialiseTape() []Symbol {
	return []Symbol{One, One, Blank, Blank}
}

// loads the program
func getProgram() map[Key]Transition {
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
	return transitions
}
