package app

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/vadim-dmitriev/sittme/stream"
	"github.com/valyala/fasthttp"
)

type Service struct {
	Handler fasthttp.RequestHandler
	Streams []stream.Stream
}

func New() *Service {
	srv := &Service{
		Streams: make([]stream.Stream, 0),
	}

	srv.createHandler()

	return srv
}

func (srv *Service) createNewStream() stream.Stream {
	newStream := stream.New()

	srv.Streams = append(srv.Streams, newStream)

	return newStream
}

func (srv *Service) getStreams() []stream.Stream {
	return srv.Streams
}

func (srv *Service) deleteStream(streamUUID uuid.UUID) error {
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
