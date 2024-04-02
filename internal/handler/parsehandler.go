package handler

import (
	"errors"
	"github.com/run-bigpig/svrw/internal/parser"
	"github.com/run-bigpig/svrw/internal/parser/douyin"
	"github.com/run-bigpig/svrw/internal/parser/pipixia"
	"github.com/run-bigpig/svrw/internal/parser/weishi"
	"github.com/valyala/fasthttp"
	"strings"
)

func ParseHandler(ctx *fasthttp.RequestCtx) {
	parseUrl := ctx.QueryArgs().Peek("url")
	if len(parseUrl) == 0 {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	p, err := loadParser(string(parseUrl))
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}
	result, err := p.Parse()
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(result.ToJson())
}

func loadParser(url string) (parser.Parser, error) {
	switch {
	case strings.Contains(url, "weishi"):
		return weishi.NewParser(url), nil
	case strings.Contains(url, "pipix"):
		return pipixia.NewParser(url), nil
	case strings.Contains(url, "douyin"):
		return douyin.NewParser(url), nil
	default:
		return nil, errors.New("not support")
	}
}
