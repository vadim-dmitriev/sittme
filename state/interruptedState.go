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
	return true
}

func (s stateInterrupted) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.stateString + `"`), nil
}
