package app

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
)

func (srv *Service) createStreamHandler() fasthttp.RequestHandler {

	return func(ctx *fasthttp.RequestCtx) {
		newStream := srv.createNewStream()

		json.NewEncoder(ctx).Encode(newStream)

		ctx.SetContentType("application/json")
		ctx.SetStatusCode(fasthttp.StatusCreated)
	}

}

func (srv *Service) deleteStreamHandler() fasthttp.RequestHandler {

	return func(ctx *fasthttp.RequestCtx) {
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

}

func (srv *Service) showStreamsHandler() fasthttp.RequestHandler {

	return func(ctx *fasthttp.RequestCtx) {
		streams := srv.getStreams()

		json.NewEncoder(ctx).Encode(streams)

		ctx.SetStatusCode(fasthttp.StatusOK)
		ctx.SetContentType("application/json")
	}

}

func (srv *Service) changeStreamStateHandler() fasthttp.RequestHandler {

	return func(ctx *fasthttp.RequestCtx) {

	}

}
