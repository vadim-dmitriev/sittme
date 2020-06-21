package state

import "fmt"

type state struct {
	stateString string
}

// NewState возвращает объект состояния на основе его текстового представления
func NewState(stateString string) (Stater, error) {

	switch stateString {

	case createdStateString:
		return NewCreated(), nil

	case activeStateString:
		return NewActive(), nil

	case interruptedStateString:
		return NewInterrupted(), nil

	case finishedStateString:
		return NewFinished(), nil

	default:
		return nil, fmt.Errorf("unsupported state '%s'", stateString)
	}

}
