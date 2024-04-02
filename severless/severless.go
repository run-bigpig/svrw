package handler

import (
	"github.com/run-bigpig/svrw/internal/parser"
	"net/http"
)

func ServerLess(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		w.Write([]byte("url is required"))
		return
	}
	p, err := parser.LoadParser(url)
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
