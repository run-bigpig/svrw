package main

import (
	"github.com/run-bigpig/svrw/internal/run"
	"github.com/run-bigpig/svrw/internal/severless"
	"net/http"
)

func main() {
	run.Run()
}

func Handler(w http.ResponseWriter, r *http.Request) {
	severless.ServerLess(w, r)
}
