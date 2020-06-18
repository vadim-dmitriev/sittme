package state

const (
	finishedStateString = "finished"
)

type StateFinished struct {
	stateString string
}

func NewFinished() Stater {
	return StateFinished{
		stateString: finishedStateString,
	}
}

func (s StateFinished) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.stateString + `"`), nil
}
