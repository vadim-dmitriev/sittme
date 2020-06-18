package app

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func (srv *Service) createHandler() {
	router := fasthttprouter.New()

	router.GET("/api/v1/streams", srv.showStreamsHandler())
	router.POST("/api/v1/streams", srv.createStreamHandler())
	router.DELETE("/api/v1/streams/:uuid", srv.deleteStreamHandler())
	router.PATCH("/api/v1/streams/:uuid", srv.changeStreamStateHandler())

	srv.handler = logMiddleware(router.Handler)
}

func (srv *Service) Handler() fasthttp.RequestHandler {
	return srv.handler
}
