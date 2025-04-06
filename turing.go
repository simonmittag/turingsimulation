package main

import "fmt"

type Symbol rune

const (
	Zero  Symbol = '0'
	One   Symbol = '1'
	Blank Symbol = '_'
)

// Direction of the tape head
const (
	Left  = -1
	Right = 1
)

// Transition represents a single Turing machine rule
type Transition struct {
	Write     Symbol
	Move      int
	NextState string
}

// Key to lookup transitions
type Key struct {
	State  string
	Symbol Symbol
}
type TuringMachine struct {
	Tape    []Symbol
	Head    int
	State   string
	Program map[Key]Transition
}

// This is the machine loop
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

func (tm *TuringMachine) MoveTheHead(trans *Transition) {
	tm.Head += trans.Move
	if tm.Head < 0 {
		tm.Head = 0 // prevent negative index
		tm.Tape = append([]Symbol{Blank}, tm.Tape...)
	} else if tm.Head >= len(tm.Tape) {
		tm.Tape = append(tm.Tape, Blank)
	}
}

func (tm *TuringMachine) WriteToTape(trans *Transition) {
	tm.Tape[tm.Head] = trans.Write
}

func (tm *TuringMachine) DetermineNextTransition(tapeValue Symbol) (*Transition, bool) {
	key := tm.HowToFindNextTransition(tapeValue)
	transition, ok := tm.Program[key]
	if !ok {
		return nil, true // if no transition found
	}
	return &transition, false
}

func (tm *TuringMachine) HowToFindNextTransition(readFromTape Symbol) Key {
	return Key{tm.State, readFromTape}
}

func (tm *TuringMachine) ReadFromTape() Symbol {
	return tm.Tape[tm.Head]
}

func (tm *TuringMachine) Run() {
	for tm.PerformSingleStep() {
		// uncomment to debug step-by-step
		fmt.Printf("%s | %c | %d | %v\n", tm.State, tm.Tape[tm.Head], tm.Head, string(tm.Tape))
	}
}
