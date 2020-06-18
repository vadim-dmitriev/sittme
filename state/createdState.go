package state

const (
	createdStateString = "created"
)

type StateCreated struct {
	state
}

func NewCreated() Stater {
	return StateCreated{
		state{
			createdStateString,
		},
	}
}

func (s StateCreated) ChangeTo(newState Stater) (Stater, error) {

	return nil, nil
}

func (s StateCreated) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.stateString + `"`), nil
}
