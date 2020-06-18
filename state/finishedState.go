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

func (s stateFinished) IsAllowChangeTo(newState Stater) bool {
	// Из состояние Finished никуда нельзя перейти
	return false
}

func (s stateFinished) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.stateString + `"`), nil
}
