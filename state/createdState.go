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
	switch newState.(type) {

	case StateActive:
		return true

	default:
		return false
	}

}

func (s StateCreated) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.stateString + `"`), nil
}

func (s StateCreated) String() string {
	return s.stateString
}
