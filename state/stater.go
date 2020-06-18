package state

type Stater interface {
	IsAllowChangeTo(Stater) bool
	MarshalJSON() ([]byte, error)
}
