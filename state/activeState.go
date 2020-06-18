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

func (s stateActive) ChangeTo(newState Stater) (Stater, error) {

	return nil, nil
}

func (s stateActive) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.stateString + `"`), nil
}
