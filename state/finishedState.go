package state

const (
	finishedStateString = "finished"
)

// Finished структура состояния 'завершен'.
// Имплементирует интерфейс Stater
type Finished struct {
	state
}

// NewFinished создает новый объект структуры Finished
func NewFinished() Stater {
	return Finished{
		state{
			finishedStateString,
		},
	}
}

// IsAllowChangeTo проверяет возможность перехода из текущего
// состояния в новое
func (s Finished) IsAllowChangeTo(newState Stater) bool {
	// Из состояние Finished никуда нельзя перейти
	return false
}

// MarshalJSON необходим для имплементации интерфейса JSONMarshaller
func (s Finished) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.stateString + `"`), nil
}

// String нужен для имплементации интерфейса Stringer
func (s Finished) String() string {
	return s.stateString
}
