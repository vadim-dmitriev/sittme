package app

import (
	"time"

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
		return streamNotFoundError{uuid}
	}

	return srv.db.Delete(uuid)
}

func (srv *Service) setNewState(uuid uuid.UUID, newString string) error {
	stream, err := srv.db.Select(uuid)
	if err != nil {
		return streamNotFoundError{uuid}
	}

	newState, err := state.NewState(newString)
	if err != nil {
		return err
	}

	currentState := stream.GetState()

	if newState == currentState {
		return nil
	}

	if !currentState.IsAllowChangeTo(newState) {
		return canNotChangeStateError{
			uuid,
			currentState,
			newState,
		}
	}

	newStreamChan := make(chan state.Stater)

	go func(ch chan state.Stater) {
		timer := time.NewTimer(2 * time.Second)

		select {
		case newState := <-ch:
			if currentState.IsAllowChangeTo(newState) {
				srv.db.Update(uuid, state.NewActive())
			}

		case <-timer.C:
			srv.db.Update(uuid, state.NewFinished())
		}

	}(newStreamChan)
	newStreamChan <- newState

	return srv.db.Update(uuid, newState)
}
