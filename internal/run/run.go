package run

import (
	"flag"
	"github.com/run-bigpig/svrw/internal/config"
	"github.com/run-bigpig/svrw/internal/router"
	"github.com/valyala/fasthttp"
	"log"
)

func Run() {
	var c config.Config
	flagConfig(&c)
	app := router.NewRouter()
	log.Printf("listen on %s\n", c.Address)
	err := fasthttp.ListenAndServe(c.Address, app.Handler)
	if err != nil {
		panic(err)
	}
}

func flagConfig(c *config.Config) {
	address := flag.String("addr", ":10800", "listen address")
	biliBiliCookie := flag.String("bc", "", "bilibili cookie")
	flag.Parse()
	c.Address = *address
	c.BiliBiliCookie = *biliBiliCookie
	config.Init(c)
}
