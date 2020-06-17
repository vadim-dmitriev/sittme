package app

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/vadim-dmitriev/sittme/stream"
	"github.com/valyala/fasthttp"
)

type Service struct {
	Streams []stream.Stream

	sync.RWMutex
	handler fasthttp.RequestHandler
}

func New() *Service {
	srv := &Service{
		make([]stream.Stream, 0),
		sync.RWMutex{},
		nil,
	}

	srv.createHandler()

	return srv
}

func (srv *Service) createNewStream() stream.Stream {
	srv.Lock()
	defer srv.Unlock()

	newStream := stream.New()

	srv.Streams = append(srv.Streams, newStream)

	return newStream
}

func (srv *Service) getStreams() []stream.Stream {
	srv.RLock()
	defer srv.RUnlock()

	return srv.Streams
}

func (srv *Service) deleteStream(streamUUID uuid.UUID) error {
	srv.Lock()
	defer srv.Unlock()

	for i, stream := range srv.Streams {
		if stream.UID == streamUUID {
			srv.Streams[i] = srv.Streams[len(srv.Streams)-1]
			srv.Streams = srv.Streams[:len(srv.Streams)-1]
			return nil
		}
	}

	// элемент не был найден и, соответственно, удален
	return fmt.Errorf("stream %s not found", streamUUID.String())
}
