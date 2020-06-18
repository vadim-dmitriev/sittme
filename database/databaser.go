package database

import (
	"github.com/google/uuid"
	"github.com/vadim-dmitriev/sittme/state"
	"github.com/vadim-dmitriev/sittme/stream"
)

type Databaser interface {
	Insert(stream.Stream) error
	Select(uuid uuid.UUID) (stream.Stream, error)
	SelectAll() []stream.Stream
	Delete(uuid.UUID) error
	Update(uuid.UUID, state.Stater) error
}
