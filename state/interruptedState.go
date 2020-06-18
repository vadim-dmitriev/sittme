package state

const (
	interruptedStateString = "interrupted"
)

type stateInterrupted struct {
	state
}

func newInterrupted() Stater {
	return stateInterrupted{
		state{
			interruptedStateString,
		},
	}
}

func (s stateInterrupted) IsAllowChangeTo(newState Stater) bool {

	// Из состояния Interrupted можно перейти в состояние
	// Finished ИЛИ вернуться в состояние Active
	switch newState.(type) {
	case stateActive, stateFinished:
		return true

	default:
		return false
	}

}

func (s stateInterrupted) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.stateString + `"`), nil
}
