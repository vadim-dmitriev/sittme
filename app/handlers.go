package app

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
)

type response struct {
	Data  interface{} `json:"data"`
	Error interface{} `json:"error"`
}

func (srv *Service) createStreamHandler() fasthttp.RequestHandler {

	return func(ctx *fasthttp.RequestCtx) {
		newStream, err := srv.createNewStream()
		if err != nil {
			serializeResponse(ctx, err, fasthttp.StatusInternalServerError)
		}

		serializeResponse(ctx, newStream, fasthttp.StatusCreated)
	}

}

func (srv *Service) deleteStreamHandler() fasthttp.RequestHandler {

	return func(ctx *fasthttp.RequestCtx) {

		uuidString := ctx.UserValue("uuid").(string)

		uuid, err := uuid.Parse(uuidString)
		if err != nil {
			serializeResponse(ctx, err, fasthttp.StatusBadRequest)
			return
		}

		if err := srv.deleteStream(uuid); err != nil {
			serializeResponse(ctx, err, fasthttp.StatusNotFound)
			return
		}

		serializeResponse(ctx, nil, fasthttp.StatusOK)
	}

}

func (srv *Service) showStreamsHandler() fasthttp.RequestHandler {

	return func(ctx *fasthttp.RequestCtx) {
		streams, err := srv.getStreams()
		if err != nil {
			serializeResponse(ctx, err, fasthttp.StatusInternalServerError)
			return
		}

		serializeResponse(ctx, streams, fasthttp.StatusOK)
	}

}

func (srv *Service) changeStreamStateHandler() fasthttp.RequestHandler {
	payload := struct {
		NewState string `json:"state"`
	}{}

	return func(ctx *fasthttp.RequestCtx) {
		if err := json.Unmarshal(ctx.PostBody(), &payload); err != nil || payload.NewState == "" {
			serializeResponse(ctx, err, fasthttp.StatusBadRequest)
			return
		}

		uuidString := ctx.UserValue("uuid").(string)
		uuid, err := uuid.Parse(uuidString)
		if err != nil {
			serializeResponse(ctx, err, fasthttp.StatusBadRequest)
			return
		}

		if err := srv.setNewState(uuid, payload.NewState); err != nil {

			switch err.(type) {
			case streamNotFoundError:
				serializeResponse(ctx, err, fasthttp.StatusNotFound)

			case canNotChangeStateError:
				serializeResponse(ctx, err, fasthttp.StatusUnprocessableEntity)

			default:
				serializeResponse(ctx, err, fasthttp.StatusInternalServerError)
			}

			return
		}

		serializeResponse(ctx, nil, fasthttp.StatusNoContent)
	}

}

func serializeResponse(ctx *fasthttp.RequestCtx, data interface{}, statusCode int) {
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(statusCode)

	encoder := json.NewEncoder(ctx)

	switch data := data.(type) {
	case error:
		encoder.Encode(
			response{
				Data:  nil,
				Error: data.Error(),
			},
		)

	default:
		if data == nil {
			return
		}
		encoder.Encode(
			response{
				Data:  data,
				Error: nil,
			},
		)

	}
}
