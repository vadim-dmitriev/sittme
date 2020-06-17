package state

type StateCreated struct {
	stateString string
}

func NewCreated() Stater {
	return StateCreated{
		stateString: createdStateString,
	}
}

func (s StateCreated) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.stateString + `"`), nil
}
