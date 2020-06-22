package app

import (
	"time"

	"github.com/vadim-dmitriev/sittme/common"

	"github.com/google/uuid"
	"github.com/vadim-dmitriev/sittme/database"
	"github.com/vadim-dmitriev/sittme/state"
	"github.com/vadim-dmitriev/sittme/stream"
	"github.com/valyala/fasthttp"
)

// Service структура сервиса. Является связующей между
// web, хранилищем и бизнес логикой.
type Service struct {
	db      database.Databaser
	handler fasthttp.RequestHandler
	config  *common.Config
}

// New создает новый объект структуры Service
func New(cfg *common.Config) *Service {
	srv := &Service{
		db:      database.NewInMemory(),
		handler: nil,
		config:  cfg,
	}

	srv.createHandler()

	return srv
}

func (srv *Service) createNewStream() (stream.Stream, error) {
	newStream := stream.New()

	err := srv.db.Insert(newStream)

	return newStream, err
}

func (srv *Service) getStreams() ([]stream.Stream, error) {
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
	selectedStream, err := srv.db.Select(uuid)
	if err != nil {
		return streamNotFoundError{uuid}
	}

	newState, err := state.NewState(newString)
	if err != nil {
		return err
	}

	currentState := selectedStream.GetState()

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

	go func(s stream.Stream) {
		currentState := <-s.StateChan
		if currentState == nil {
			return
		}

		if !currentState.IsAllowChangeTo(state.NewFinished()) {
			return
		}

		timer := time.NewTimer(srv.config.Timeout)
		select {
		case newState := <-s.StateChan:
			if newState != nil && currentState.IsAllowChangeTo(newState) {
				srv.db.Update(s.UUID, newState)
			}

		case <-timer.C:
			srv.db.Update(s.UUID, state.NewFinished())

		}

	}(selectedStream)

	selectedStream.StateChan <- newState

	return srv.db.Update(uuid, newState)
}
