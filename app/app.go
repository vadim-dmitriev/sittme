package app

import (
	"fmt"

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
	_, err := srv.db.Select(uuid)
	if err != nil {
		return fmt.Errorf("stream %s not found", uuid.String())
	}

	return srv.db.Delete(uuid)
}

func (srv *Service) setNewState(uuid uuid.UUID, newString string) error {
	stream, err := srv.db.Select(uuid)
	if err != nil {
		return fmt.Errorf("stream %s not found", uuid.String())
	}

	newState, err := state.NewState(newString)
	if err != nil {
		return err
	}

	currentState := stream.GetState()
	if !currentState.IsAllowChangeTo(newState) {
		return fmt.Errorf("can`t change state")
	}

	return srv.db.Update(uuid, newState)
}
