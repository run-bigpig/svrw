package parser

import (
	"encoding/json"
	"errors"
	"github.com/run-bigpig/svrw/internal/parser/douyin"
	"github.com/run-bigpig/svrw/internal/parser/pipixia"
	"github.com/run-bigpig/svrw/internal/parser/weishi"
	"strings"
)

type Data struct {
	Author string `json:"author"`
	Avatar string `json:"avatar"`
	Time   string `json:"time"`
	Title  string `json:"title"`
	Cover  string `json:"cover"`
	Url    string `json:"url"`
}

type ParseResult struct {
	*Data
}

type Parser interface {
	Parse() (*ParseResult, error)
}

func (pr *ParseResult) ToJson() []byte {
	jsonData, err := json.Marshal(pr)
	if err != nil {
		return nil
	}
	return jsonData
}

func LoadParser(url string) (Parser, error) {
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
