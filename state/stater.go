package state

// Stater интерфейс объединяющий все объекты состояний
type Stater interface {

	// IsAllowChangeTo проверяет возможность перехода из текущего
	// состояния в новое
	IsAllowChangeTo(Stater) bool

	// MarshalJSON необходим для имплементации интерфейса JSONMarshaller
	MarshalJSON() ([]byte, error)

	// String нужен для имплементации интерфейса Stringer
	String() string
}
