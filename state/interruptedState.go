package state

const (
	interruptedStateString = "interrupted"
)

// Interrupted структура состояния 'прерван'.
// Имплементирует интерфейс Stater
type Interrupted struct {
	state
}

// NewInterrupted создает новый объект структуры Interrupted
func NewInterrupted() Stater {
	return Interrupted{
		state{
			interruptedStateString,
		},
	}
}

// IsAllowChangeTo проверяет возможность перехода из текущего
// состояния в новое
func (s Interrupted) IsAllowChangeTo(newState Stater) bool {

	// Из состояния Interrupted можно перейти в состояние
	// Finished ИЛИ вернуться в состояние Active
	switch newState.(type) {
	case Active, Finished:
		return true

	default:
		return false
	}

}

// MarshalJSON необходим для имплементации интерфейса JSONMarshaller
func (s Interrupted) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.stateString + `"`), nil
}

// String нужен для имплементации интерфейса Stringer
func (s Interrupted) String() string {
	return s.stateString
}
