package state

const (
	activeStateString = "active"
)

type StateActive struct {
	state
}

func NewActive() Stater {
	return StateActive{
		state{
			activeStateString,
		},
	}
}

func (s StateActive) IsAllowChangeTo(newState Stater) bool {

	// Из состояния Active можно перейти только в состояние Interrupted
	switch newState.(type) {
	case StateInterrupted:
		return true

	default:
		return false
	}

}

func (s StateActive) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.stateString + `"`), nil
}
