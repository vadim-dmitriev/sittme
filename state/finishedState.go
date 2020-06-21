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

func (s StateFinished) IsAllowChangeTo(newState Stater) bool {
	// Из состояние Finished никуда нельзя перейти
	return false
}

func (s StateFinished) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.stateString + `"`), nil
}

func (s StateFinished) String() string {
	return s.stateString
}
