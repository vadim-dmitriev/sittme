package state

const (
	interruptedStateString = "interrupted"
)

type StateInterrupted struct {
	stateString string
}

func NewInterrupted() Stater {
	return StateInterrupted{
		stateString: interruptedStateString,
	}
}

func (s StateInterrupted) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.stateString + `"`), nil
}
