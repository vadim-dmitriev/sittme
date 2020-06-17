package main

import (
	"encoding/json"

	"github.com/google/uuid"

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
	ctx.SetContentType("application/json")

	uidString := ctx.UserValue("uid").(string)

	uid, err := uuid.Parse(uidString)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	if err := srv.deleteStream(uid); err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	ctx.SetStatusCode(fasthttp.StatusOK)
}

func (srv *service) showStreamsHandler(ctx *fasthttp.RequestCtx) {
	streams := srv.getStreams()

	json.NewEncoder(ctx).Encode(streams)

	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("application/json")
}

func (srv *service) changeStreamStateHandler(ctx *fasthttp.RequestCtx) {

}
