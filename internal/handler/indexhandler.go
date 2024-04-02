package handler

import "github.com/valyala/fasthttp"

func IndexHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/plain; charset=utf-8")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.WriteString("Hello, World!")
}
