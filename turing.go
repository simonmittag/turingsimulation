package main

import "fmt"

// Symbol represents a single character or rune used in a Turing machine's tape or transitions.
type Symbol rune

// Zero represents the symbol '0'.
// One represents the symbol '1'.
// Blank represents the blank symbol '_'.
const (
	Zero  Symbol = '0'
	One   Symbol = '1'
	Blank Symbol = '_'
)

// Left represents the left direction with a value of -1.
// Right represents the right direction with a value of 1.
const (
	Left  = -1
	Right = 1
)

// Transition represents a Turing machine state transition.
// It specifies the symbol to write, the head movement, and the next state.
type Transition struct {
	Write     Symbol
	Move      int
	NextState string
}

// Key represents a combination of the current state and a symbol used to determine the next transition in a Turing Machine.
type Key struct {
	State  string
	Symbol Symbol
}

// TuringMachine represents a Turing Machine model with a tape, head position, state, and program transitions.
type TuringMachine struct {
	Tape    []Symbol
	Head    int
	State   string
	Program map[Key]Transition
}

// PerformSingleStep executes a single transition in the Turing Machine based on the current state and tape value.
func (tm *TuringMachine) PerformSingleStep() bool {
	// 1. read the tape value from the current position
	tapeValue := tm.ReadFromTape()

	// 2. figure out next transition in the program
	trans, done := tm.DetermineNextTransition(tapeValue)
	if done {
		//if there isn't one we stop the execution
		return false
	}

	// 3. execute the transition
	tm.WriteToTape(trans)
	tm.MoveTheHead(trans)
	tm.State = trans.NextState
	return true
}

// MoveTheHead adjusts the machine's head position based on the given transition and extends the tape as needed.
func (tm *TuringMachine) MoveTheHead(trans *Transition) {
	tm.Head += trans.Move
	if tm.Head < 0 {
		tm.Head = 0 // prevent negative index
		tm.Tape = append([]Symbol{Blank}, tm.Tape...)
	} else if tm.Head >= len(tm.Tape) {
		tm.Tape = append(tm.Tape, Blank)
	}
}

// WriteToTape writes the specified symbol from the transition to the current position on the tape.
func (tm *TuringMachine) WriteToTape(trans *Transition) {
	tm.Tape[tm.Head] = trans.Write
}

// DetermineNextTransition determines the next state transition based on the current tape symbol.
// Returns the transition and a flag indicating if no valid transition is found.
func (tm *TuringMachine) DetermineNextTransition(tapeValue Symbol) (*Transition, bool) {
	key := tm.HowToFindNextTransition(tapeValue)
	transition, ok := tm.Program[key]
	if !ok {
		return nil, true // if no transition found
	}
	return &transition, false
}

// HowToFindNextTransition generates a unique Key using the Turing Machine's current state and the given tape symbol.
func (tm *TuringMachine) HowToFindNextTransition(readFromTape Symbol) Key {
	return Key{tm.State, readFromTape}
}

// ReadFromTape retrieves the current symbol at the tape position indicated by the Turing machine's head.
func (tm *TuringMachine) ReadFromTape() Symbol {
	return tm.Tape[tm.Head]
}

// Run executes the Turing machine starting from the initial state and continues until no further transitions are possible.
func (tm *TuringMachine) Run() {
	for tm.PerformSingleStep() {
		//log what the machine is doing
		fmt.Printf("%s | %c | %d | %v\n", tm.State, tm.Tape[tm.Head], tm.Head, string(tm.Tape))
	}
}
