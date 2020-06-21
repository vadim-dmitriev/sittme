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

	StateChan chan state.Stater `json:"-"`
}

// New cоздает новый объект трансляции, у которого
// состояние 'created'
func New() Stream {
	stream := Stream{
		UUID:      generateUUID(),
		StateChan: make(chan state.Stater),
	}

	stream.Attributes.DateModified = time.Now()
	stream.Attributes.State = state.NewCreated()

	return stream
}

func (s *Stream) GetState() state.Stater {
	return s.Attributes.State
}

func (s *Stream) SetState(newState state.Stater) {
	s.Attributes.State = newState
}

func generateUUID() uuid.UUID {
	uuid, _ := uuid.NewRandom()
	return uuid
}
