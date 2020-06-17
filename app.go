package main

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
