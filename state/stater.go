package state

type Stater interface {
	MarshalJSON() ([]byte, error)
}
