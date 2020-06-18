package state

const (
	finishedStateString = "finished"
)

type stateFinished struct {
	state
}

func newFinished() Stater {
	return stateFinished{
		state{
			finishedStateString,
		},
	}
}

func (s stateFinished) ChangeTo(newState Stater) (Stater, error) {
	return nil, nil
}

func (s stateFinished) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.stateString + `"`), nil
}
