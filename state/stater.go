package state

type Stater interface {
	ChangeTo(Stater) (Stater, error)
	MarshalJSON() ([]byte, error)
}
