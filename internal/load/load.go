package load

import (
	"errors"
	"github.com/run-bigpig/svrw/internal/parser"
	"github.com/run-bigpig/svrw/internal/parser/douyin"
	"github.com/run-bigpig/svrw/internal/parser/pipixia"
	"github.com/run-bigpig/svrw/internal/parser/weishi"
	"github.com/run-bigpig/svrw/internal/parser/zuiyou"
	"strings"
)

func LoadParser(url string) (parser.Parser, error) {
	switch {
	case strings.Contains(url, "weishi"):
		return weishi.NewParser(url), nil
	case strings.Contains(url, "pipix"):
		return pipixia.NewParser(url), nil
	case strings.Contains(url, "douyin"):
		return douyin.NewParser(url), nil
	case strings.Contains(url, "zuiyou") || strings.Contains(url, "xiaochuankeji"):
		return zuiyou.NewParser(url), nil
	default:
		return nil, errors.New("not support")
	}
}
