package stream

import (
	"time"

	"github.com/google/uuid"
	"github.com/vadim-dmitriev/sittme/state"
)

type Stream struct {
	UUID       uuid.UUID `json:"id"`
	Attributes struct {
		State        state.Stater `json:"state"`
		DateModified time.Time    `json:"date_modified"`
	} `json:"attributes"`
}

// New cоздает новый объект трансляции, у которого
// состояние 'created'
func New() *Stream {
	stream := Stream{
		UUID: generateUUID(),
	}

	stream.Attributes.DateModified = time.Now()
	stream.Attributes.State = state.NewCreated()

	return &stream
}

func (s *Stream) SetNewState(newState *state.Stater) error {

	// вызов stater.ChangeTo()

	return nil
}

func generateUUID() uuid.UUID {
	uuid, _ := uuid.NewRandom()
	return uuid
}
