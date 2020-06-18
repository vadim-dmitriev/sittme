package state

import "fmt"

/*
	TODO: Тут можно определить тип state, от которого будут наследоваться
	конкретные типы состояний

	type state struct {
		stateString string
	}

*/
func NewStater(stateString string) (Stater, error) {

	switch stateString {

	case createdStateString:
		return NewCreated(), nil

	case activeStateString:
		return NewActive(), nil

	default:
		return nil, fmt.Errorf("unsupported state '%s'", stateString)
	}

}
