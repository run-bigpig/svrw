package handler

import (
	"errors"
	"github.com/run-bigpig/svrw/internal/parser"
	"github.com/run-bigpig/svrw/internal/parser/douyin"
	"github.com/run-bigpig/svrw/internal/parser/pipixia"
	"github.com/run-bigpig/svrw/internal/parser/weishi"
	"github.com/run-bigpig/svrw/internal/response"
	"github.com/valyala/fasthttp"
	"strings"
)

func ParseHandler(ctx *fasthttp.RequestCtx) {
	parseUrl := ctx.QueryArgs().Peek("url")
	if len(parseUrl) == 0 {
		response.Error(ctx, errors.New("url is required"))
		return
	}
	p, err := loadParser(string(parseUrl))
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
