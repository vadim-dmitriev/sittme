package state

const (
	interruptedStateString = "interrupted"
)

type StateInterrupted struct {
	state
}

func NewInterrupted() Stater {
	return StateInterrupted{
		state{
			interruptedStateString,
		},
	}
}

func (s StateInterrupted) IsAllowChangeTo(newState Stater) bool {

	// Из состояния Interrupted можно перейти в состояние
	// Finished ИЛИ вернуться в состояние Active
	switch newState.(type) {
	case StateActive, StateFinished:
		return true

	default:
		return false
	}

}

func (s StateInterrupted) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.stateString + `"`), nil
}

func (s StateInterrupted) String() string {
	return s.stateString
}
