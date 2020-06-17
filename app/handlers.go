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

		uuidString := ctx.UserValue("uuid").(string)

		uuid, err := uuid.Parse(uuidString)
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			return
		}

		if err := srv.deleteStream(uuid); err != nil {
			ctx.SetStatusCode(fasthttp.StatusNotFound)
			return
		}

		ctx.SetStatusCode(fasthttp.StatusNoContent)
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
	payload := struct {
		NewState string `json:"state"`
	}{}

	return func(ctx *fasthttp.RequestCtx) {
		ctx.SetContentType("application/json")

		if err := json.Unmarshal(ctx.PostBody(), &payload); err != nil || payload.NewState == "" {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			return
		}

		uuidString := ctx.UserValue("uuid").(string)
		uuid, err := uuid.Parse(uuidString)
		if err != nil {
			ctx.SetStatusCode(fasthttp.StatusBadRequest)
			return
		}

		if err := srv.setNewState(uuid, payload.NewState); err != nil {
			// изменение состояние противоречит условиям
			ctx.SetStatusCode(fasthttp.StatusUnprocessableEntity)
			return
		}
	}

}
