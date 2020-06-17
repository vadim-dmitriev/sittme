package stream

import (
	"bytes"
	"time"

	"github.com/google/uuid"
)

type state int

const (
	createdState = iota
	activeState
	interruptedState
	finishedState
)

var (
	iotaStringRepresentation = map[state]string{
		createdState:     "created",
		activeState:      "active",
		interruptedState: "interrupted",
		finishedState:    "finished",
	}
)

func (s state) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(iotaStringRepresentation[s])
	buffer.WriteString(`"`)

	return buffer.Bytes(), nil
}

type Stream struct {
	UID        uuid.UUID `json:"id"`
	Attributes struct {
		State        state     `json:"state"`
		DateModified time.Time `json:"date_modified"`
	} `json:"attributes"`
}

// New cоздает новый объект трансляции, у которого
// состояние Created
func New() Stream {
	stream := Stream{
		UID: generateUUID(),
	}

	stream.Attributes.DateModified = time.Now()
	stream.Attributes.State = createdState

	return stream
}

func generateUUID() uuid.UUID {
	uuid, _ := uuid.NewRandom()
	return uuid
}
