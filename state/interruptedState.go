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

func (s StateInterrupted) ChangeTo(newState Stater) (Stater, error) {
	return nil, nil
}

func (s StateInterrupted) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.stateString + `"`), nil
}
