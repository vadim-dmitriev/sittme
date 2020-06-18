package state

const (
	finishedStateString = "finished"
)

type StateFinished struct {
	state
}

func NewFinished() Stater {
	return StateFinished{
		state{
			finishedStateString,
		},
	}
}

func (s StateFinished) ChangeTo(newState Stater) (Stater, error) {
	return nil, nil
}

func (s StateFinished) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.stateString + `"`), nil
}
