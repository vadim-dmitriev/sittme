package state

import "fmt"

type state struct {
	stateString string
}

func NewState(stateString string) (Stater, error) {

	switch stateString {

	case createdStateString:
		return NewCreated(), nil

	case activeStateString:
		return newActive(), nil

	case interruptedStateString:
		return newInterrupted(), nil

	case finishedStateString:
		return newFinished(), nil

	default:
		return nil, fmt.Errorf("unsupported state '%s'", stateString)
	}

}
