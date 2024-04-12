package load

import (
	"errors"
	"github.com/run-bigpig/svrw/internal/parser"
	"github.com/run-bigpig/svrw/internal/parser/bilibili"
	"github.com/run-bigpig/svrw/internal/parser/douyin"
	"github.com/run-bigpig/svrw/internal/parser/kuaishou"
	"github.com/run-bigpig/svrw/internal/parser/pipixia"
	"github.com/run-bigpig/svrw/internal/parser/weishi"
	"github.com/run-bigpig/svrw/internal/parser/xiaohongshu"
	"github.com/run-bigpig/svrw/internal/parser/xigua"
	"github.com/run-bigpig/svrw/internal/parser/zuiyou"
	"strings"
)

func LoadParser(url string) (parser.Parser, error) {
	url = strings.Trim(url, " ")
	switch {
	case strings.Contains(url, "weishi"):
		return weishi.NewParser(url), nil
	case strings.Contains(url, "pipix"):
		return pipixia.NewParser(url), nil
	case strings.Contains(url, "douyin"):
		return douyin.NewParser(url), nil
	case strings.Contains(url, "zuiyou") || strings.Contains(url, "xiaochuankeji"):
		return zuiyou.NewParser(url), nil
	case strings.Contains(url, "kuaishou"):
		return kuaishou.NewParser(url), nil
	case strings.Contains(url, "ixigua"):
		return xigua.NewParser(url), nil
	case strings.Contains(url, "xhslink.com") || strings.Contains(url, "xiaohongshu.com"):
		return xiaohongshu.NewParser(url), nil
	case strings.Contains(url, "b23.tv") || strings.Contains(url, "bilibili.com"):
		return bilibili.NewParser(url), nil
	default:
		return nil, errors.New("not support")
	}
}
