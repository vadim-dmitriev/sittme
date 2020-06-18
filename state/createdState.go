package state

const (
	createdStateString = "created"
)

type StateCreated struct {
	state
}

func NewCreated() Stater {
	return StateCreated{
		state{
			createdStateString,
		},
	}
}

func (s StateCreated) IsAllowChangeTo(newState Stater) bool {

	// Из состояния Created можно перейти только в состояние Active
	if _, ok := newState.(stateActive); !ok {
		return false
	}

	return true
}

func (s StateCreated) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.stateString + `"`), nil
}
