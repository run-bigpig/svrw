package handler

import (
	"github.com/run-bigpig/svrw/internal/load"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		w.Write([]byte("url is required"))
		return
	}
	p, err := load.LoadParser(url)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	result, err := p.Parse()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(result.ToJson())
}
