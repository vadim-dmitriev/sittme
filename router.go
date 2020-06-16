package main

import (
	"encoding/json"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func newRouter(srv *service) *fasthttprouter.Router {
	router := fasthttprouter.New()

	router.GET("/api/v1/streams", srv.showStreamsHandler)
	router.POST("/api/v1/streams", srv.createStreamHandler)
	router.DELETE("/api/v1/streams/:uid", srv.deleteStreamHandler)
	router.PATCH("/api/v1/streams/:uid", srv.changeStreamStateHandler)

	return router
}

func (srv *service) createStreamHandler(ctx *fasthttp.RequestCtx) {
	newStream := srv.createNewStream()

	json.NewEncoder(ctx).Encode(newStream)

	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusCreated)
}

func (srv *service) deleteStreamHandler(ctx *fasthttp.RequestCtx) {

}

func (srv *service) showStreamsHandler(ctx *fasthttp.RequestCtx) {

}

func (srv *service) changeStreamStateHandler(ctx *fasthttp.RequestCtx) {

}
