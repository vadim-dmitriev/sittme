package stream

import (
	"time"

	"github.com/google/uuid"
	"github.com/vadim-dmitriev/sittme/state"
)

// Stream структура, описывающая сущность Трансляция
type Stream struct {
	// UUID уникальный идентификатор трансляции.
	// Представляет из себя UUIDv4
	UUID uuid.UUID `json:"id"`

	// Attributes описывает свойства трансляции
	Attributes struct {

		// State содержит текущее состояние трансляции
		State state.Stater `json:"state"`

		// DateModified содержит время в формате RFC 3339
		DateModified time.Time `json:"date_modified"`
	} `json:"attributes"`

	// StateChan необходим, чтобы связать объект трансляции
	// с горутиной, в которой будет происходить отложеный
	// переход между состояниями
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

// GetState возвращает текущее состояние трансляции
func (s *Stream) GetState() state.Stater {
	return s.Attributes.State
}

// SetState устанавливает новое состояние и устанавливает
// текущее время в поле DateModified
func (s *Stream) SetState(newState state.Stater) {
	s.Attributes.DateModified = time.Now()
	s.Attributes.State = newState
}

func generateUUID() uuid.UUID {
	uuid, _ := uuid.NewRandom()
	return uuid
}
