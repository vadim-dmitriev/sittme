package state

const (
	activeStateString = "active"
)

type StateActive struct {
	state
}

func NewActive() Stater {
	return StateActive{
		state{
			activeStateString,
		},
	}
}

func (s StateActive) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.stateString + `"`), nil
}
