package main

import (
	"strings"
)

// (integer|decimal) (exponent)?
// integer := (sign)? (digits)
// decimal := (sign)? (digits dot | (digits)? dot digits)
// exponent := (e|E) (integer)

type State int

const (
	StateInit State = 0
	// StateNumber can be an integer or decimal, until a token is "."
	StateSign   = 1
	StateNumber = 2
	// State Dot is to differentiate between 0.9 and .9
	StateDot     = 3
	StateDecimal = 4

	StateExponent = 10

	StateTerminal = 20
)

var (
	stateTransitions map[State]map[State]bool = map[State]map[State]bool{
		// StateInit can be transition to anything
		StateInit: {
			StateSign:   true,
			StateNumber: true,
			StateDot:    true,
		},
		StateSign: {
			StateNumber: true,
			StateDot:    true,
		},
		StateNumber: {
			StateNumber:   true,
			StateDecimal:  true,
			StateExponent: true,
			StateTerminal: true,
		},
		StateDot: {
			StateDecimal: true,
		},
		StateDecimal: {
			StateDecimal:  true,
			StateExponent: true,
			StateTerminal: true,
		},
		StateExponent: {
			// This is managed in exponentStateTransitions
		},
		StateTerminal: {},
	}

	exponentStateTransitions map[State]map[State]bool = map[State]map[State]bool{
		// StateInit can be transition to anything
		StateExponent: {
			StateSign:   true,
			StateNumber: true,
		},
		StateSign: {
			StateNumber: true,
		},
		StateNumber: {
			StateNumber:   true,
			StateTerminal: true,
		},
		StateTerminal: {},
	}
)

func isValidStateForNumbers(state State, nextState State) bool {
	_, ok := stateTransitions[state][nextState]
	return ok
}

func isValidStateExponent(state State, nextState State) bool {
	if _, ok := exponentStateTransitions[state]; !ok {
		return false
	}
	_, ok := exponentStateTransitions[state][nextState]
	return ok
}

func isNumber(s string) bool {
	return isNumber2(s)
}

// DSA
func isNumber2(s string) bool {
	isValidState := isValidStateForNumbers
	state := StateInit

	for _, ch := range s {
		var nextState State
		nextStateTransition := isValidState
		switch ch {
		case '+', '-':
			nextState = StateSign
		case 'e', 'E':
			nextState = StateExponent
			nextStateTransition = isValidStateExponent
		case '.':
			if state == StateNumber {
				nextState = StateDecimal
			} else {
				nextState = StateDot
			}
		default:
			if !('0' <= ch && ch <= '9') {
				return false
			}

			// validate a state
			if state == StateDot || state == StateDecimal {
				nextState = StateDecimal
			} else {
				nextState = StateNumber
			}
		}

		if !isValidState(state, nextState) {
			return false
		}
		state = nextState
		isValidState = nextStateTransition
	}

	return isValidStateForNumbers(state, StateTerminal)
}

// Time: O(n), Space: O(1)
func isNumber1(s string) bool {
	if strings.ContainsAny(s, "eE") {
		// exponent validate
		substrs := strings.SplitN(strings.ToLower(s), "e", 2)
		if len(substrs[1]) == 0 {
			return false
		}
		if !isInteger(substrs[1]) {
			return false
		}

		s = substrs[0]
	}
	if len(s) == 0 {
		return false
	}

	return isInteger(s) || isDecimal(s)
}

func hasSign(substr string) bool {
	return substr[0] == '+' || substr[0] == '-'
}

func parseDigits(substr string, index int) (string, int) {
	value := ""
	for ; index < len(substr); index++ {
		if !('0' <= substr[index] && substr[index] <= '9') {
			break
		}
		value = value + string(substr[index])
	}
	return value, index
}

func isInteger(substr string) bool {
	index := 0
	if hasSign(substr) {
		index++
	}
	value, endIndex := parseDigits(substr, index)
	if value == "" {
		return false
	}

	return len(substr) == endIndex
}

func isDecimal(substr string) bool {
	index := 0
	if hasSign(substr) {
		index++
	}
	integer, index := parseDigits(substr, index)

	if len(substr) <= index {
		return false
	}
	if substr[index] != '.' {
		return false
	}

	decimal, endIndex := parseDigits(substr, index+1)
	if integer == "" && decimal == "" {
		return false
	}

	return len(substr) == endIndex
}
