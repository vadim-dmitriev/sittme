package state

const (
	createdStateString = "created"
)

// Created структура состояния 'создан'.
// Имплементирует интерфейс Stater
type Created struct {
	state
}

// NewCreated создает новый объект структуры Created
func NewCreated() Stater {
	return Created{
		state{
			createdStateString,
		},
	}
}

// IsAllowChangeTo проверяет возможность перехода из текущего
// состояния в новое
func (s Created) IsAllowChangeTo(newState Stater) bool {

	// Из состояния Created можно перейти только в состояние Active
	switch newState.(type) {

	case Active:
		return true

	default:
		return false
	}

}

// MarshalJSON необходим для имплементации интерфейса JSONMarshaller
func (s Created) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.stateString + `"`), nil
}

// String нужен для имплементации интерфейса Stringer
func (s Created) String() string {
	return s.stateString
}
