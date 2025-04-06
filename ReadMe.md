# Turing Machine Simulation in Go

## Overview

This project is a simple implementation of a **Turing Machine** written in Go. 
A Turing Machine is a theoretical computing device that operates on an infinite 
tape divided into cells. Each cell can hold a symbol, and the machine's "head" can 
perform operations (e.g., read, write, move left or right) based on a predefined 
set of rules (or "transitions") that form a program. 

The implemented Turing Machine in this project performs basic bitwise operations 
(e.g., XOR logic) on predefined 1-Bit inputs.

## How It Works

The Turing Machine consists of the following components:

1. **Tape**: A sequence of symbols (`0`, `1`, or a blank `_`) where computations happen.
2. **Head**: A pointer that can:
    - Read a symbol from the tape.
    - Write a new symbol on the tape.
    - Move left or right to another tape cell.
3. **State**: The current condition of the machine, determining its behavior.
4. **Program (Transitions)**: A set of rules that describe how the machine behaves in specific states and with a given tape symbol. Each rule specifies:
    - What symbol to write at the current position.
    - Whether to move the head left or right.
    - What the next state of the machine is.

The machine begins with an initial state, reads the tape symbol at the head's position, and applies the appropriate transition. It continues operating until a state is reached where no valid transition exists (this halts the machine).

## What the Program Does

In this implementation:
1. The tape is initialized with two binary inputs, e.g., `1, 1` (representing A and B).
2. The program adds two 1 Bit numbers with bitwise XOR and AND operations.
    - XOR determines the `SUM` of A and B.
    - AND determines the `CARRY` when both A and B are `1`.
3. The result is written to the tape, and the machine stops.

## How to Run

1. **Prerequisites**:
    - Install [Go](https://go.dev/dl/).

2. **Download and Run**:
   Download the project files and use the following command in the terminal:
```shell script
go run main.go
```
This will execute the Turing Machine and print the tape's state after all computations.

3. **Example Output**:
```
Starting Program. Machine State: q0 | Head Position: 1 | Full Tape: 11__
   Step performed. Machine State: qStoreA1 | Head Position: 2 | Full Tape: 11__
   Step performed. Machine State: qXOR3 | Head Position: 4 | Full Tape: 11_0
   Step performed. Machine State: qAND3 | Head Position: 5 | Full Tape: 11_01
   Step performed. Machine State: HALT | Head Position: 5 | Full Tape: 11_01
   Final Tape: 11_01
```

## Customizing the Turing Machine

### Changing the Input Tape
You can modify the function `initialiseTape` in `main.go` to adjust the starting values on the tape:
```textmate
func initialiseTape() []Symbol {
    return []Symbol{One, One, Blank, Blank} // Example initialization
}
```

### Adding/Modifying Transitions
The transitions are defined in the `getProgram` function as a map. You can add new transitions or adjust existing ones.

Example transition:
```textmate
{"q0", Zero}: {Zero, Right, "qStoreA0"}, // Define rule for state `q0` reading `0`
```

## License

This project is open-source and available to use under the [Apache License](LICENSE).