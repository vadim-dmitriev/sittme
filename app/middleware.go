package app

import "github.com/valyala/fasthttp"

/*
	logMiddleware логирует все запросы в следующем формате:

	YYYY/MM/DD HH:MM:SS durition(sec) #requestID - host<->remoteHost - method url - statusCode

	Пример:
	2020/06/18 12:47:38 0.000 #0000000100000001 - 127.0.0.1:8080<->127.0.0.1:50716 - GET http://localhost:8080/api/v1/streams - 200

*/
func logMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {

	return func(ctx *fasthttp.RequestCtx) {
		next(ctx)
		ctx.Logger().Printf("%d", ctx.Response.StatusCode())
	}

}
