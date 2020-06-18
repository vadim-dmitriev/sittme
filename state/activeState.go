package state

const (
	activeStateString = "active"
)

type stateActive struct {
	state
}

func newActive() Stater {
	return stateActive{
		state{
			activeStateString,
		},
	}
}

func (s stateActive) IsAllowChangeTo(newState Stater) bool {

	// Из состояния Active можно перейти только в состояние Interrupted
	switch newState.(type) {
	case stateInterrupted:
		return true

	default:
		return false
	}

}

func (s stateActive) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.stateString + `"`), nil
}
