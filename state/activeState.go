package state

const (
	activeStateString = "active"
)

// Active структура состояния 'активный'.
// Имплементирует интерфейс Stater
type Active struct {
	state
}

// NewActive создает новый объект структуры Active
func NewActive() Stater {
	return Active{
		state{
			activeStateString,
		},
	}
}

// IsAllowChangeTo проверяет возможность перехода из текущего
// состояния в новое
func (s Active) IsAllowChangeTo(newState Stater) bool {

	// Из состояния Active можно перейти только в состояние Interrupted
	switch newState.(type) {
	case Interrupted:
		return true

	default:
		return false
	}

}

// MarshalJSON необходим для имплементации интерфейса JSONMarshaller
func (s Active) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.stateString + `"`), nil
}

// String нужен для имплементации интерфейса Stringer
func (s Active) String() string {
	return s.stateString
}
