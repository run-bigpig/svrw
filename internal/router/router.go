package router

import (
	"github.com/fasthttp/router"
	"github.com/run-bigpig/svrw/internal/handler"
)

func NewRouter() *router.Router {
	r := router.New()
	r.GET("/", handler.IndexHandler)
	r.GET("/api", handler.ParseHandler)
	return r
}
