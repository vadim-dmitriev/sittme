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

func (s stateInterrupted) ChangeTo(newState Stater) (Stater, error) {
	return nil, nil
}

func (s stateInterrupted) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.stateString + `"`), nil
}
