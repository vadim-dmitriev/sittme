package state

type StateActive struct {
	stateString string
}

func NewActive() Stater {
	return StateActive{
		activeStateString,
	}
}

func (s StateActive) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.stateString + `"`), nil
}
