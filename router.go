package main

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func newRouter() *fasthttprouter.Router {
	router := fasthttprouter.New()

	router.GET("/api/v1/streams", showStreams)
	router.POST("/api/v1/streams", createStream)
	router.DELETE("/api/v1/streams/:uid", deleteStream)
	router.PATCH("/api/v1/streams/:uid", changeStreamState)

	return router
}

func createStream(ctx *fasthttp.RequestCtx) {

}

func deleteStream(ctx *fasthttp.RequestCtx) {

}

func showStreams(ctx *fasthttp.RequestCtx) {

}

func changeStreamState(ctx *fasthttp.RequestCtx) {

}
