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

	return true
}

func (s stateActive) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.stateString + `"`), nil
}
