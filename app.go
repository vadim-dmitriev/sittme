package main

import (
	"fmt"

	"github.com/google/uuid"
)

type service struct {
	streams []stream
}

func (srv *service) createNewStream() stream {
	newStream := newStream()

	srv.streams = append(srv.streams, newStream)

	return newStream
}

func (srv *service) getStreams() []stream {
	return srv.streams
}

func (srv *service) deleteStream(streamUUID uuid.UUID) error {
	for i, stream := range srv.streams {
		if stream.UID == streamUUID {
			srv.streams[i] = srv.streams[len(srv.streams)-1]
			srv.streams = srv.streams[:len(srv.streams)-1]
			return nil
		}
	}

	// элемент не был найден и, соответственно, удален
	return fmt.Errorf("stream %s not found", streamUUID.String())
}
