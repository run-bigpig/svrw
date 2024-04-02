package handler

import (
	"errors"
	"github.com/run-bigpig/svrw/internal/load"
	"github.com/run-bigpig/svrw/internal/response"
	"github.com/valyala/fasthttp"
)

func ParseHandler(ctx *fasthttp.RequestCtx) {
	parseUrl := ctx.QueryArgs().Peek("url")
	if len(parseUrl) == 0 {
		response.Error(ctx, errors.New("url is required"))
		return
	}
	p, err := load.LoadParser(string(parseUrl))
	if err != nil {
		response.Error(ctx, err)
		return
	}
	result, err := p.Parse()
	if err != nil {
		response.Error(ctx, err)
		return
	}
	response.Success(ctx, result)
}
