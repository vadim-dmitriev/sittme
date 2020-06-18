package app

import (
	"github.com/google/uuid"
	"github.com/vadim-dmitriev/sittme/database"
	"github.com/vadim-dmitriev/sittme/state"
	"github.com/vadim-dmitriev/sittme/stream"
	"github.com/valyala/fasthttp"
)

type Service struct {
	db      database.Databaser
	handler fasthttp.RequestHandler
}

func New() *Service {
	srv := &Service{
		db:      database.NewInMemory(),
		handler: nil,
	}

	srv.createHandler()

	return srv
}

func (srv *Service) createNewStream() stream.Stream {
	newStream := stream.New()

	srv.db.Insert(newStream)

	return newStream
}

func (srv *Service) getStreams() []stream.Stream {
	return srv.db.SelectAll()
}

func (srv *Service) deleteStream(uuid uuid.UUID) error {
	return srv.db.Delete(uuid)
}

func (srv *Service) setNewState(uuid uuid.UUID, newStateString string) error {
	newState, err := state.NewStater(newStateString)
	if err != nil {
		return err
	}

	return srv.db.Update(uuid, newState)
}
