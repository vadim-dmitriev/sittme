package main

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func newRouter(srv *service) *fasthttprouter.Router {
	router := fasthttprouter.New()

	router.GET("/api/v1/streams", srv.showStreams)
	router.POST("/api/v1/streams", srv.createStream)
	router.DELETE("/api/v1/streams/:uid", srv.deleteStream)
	router.PATCH("/api/v1/streams/:uid", srv.changeStreamState)

	return router
}

func (srv *service) createStream(ctx *fasthttp.RequestCtx) {

}

func (srv *service) deleteStream(ctx *fasthttp.RequestCtx) {

}

func (srv *service) showStreams(ctx *fasthttp.RequestCtx) {

}

func (srv *service) changeStreamState(ctx *fasthttp.RequestCtx) {

}
