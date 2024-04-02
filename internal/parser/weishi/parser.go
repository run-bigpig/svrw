package weishi

import (
	"encoding/json"
	"errors"
	"github.com/dop251/goja"
	"github.com/run-bigpig/svrw/internal/consts"
	"github.com/run-bigpig/svrw/internal/parser"
	"github.com/run-bigpig/svrw/internal/utils"
	"log"
	"regexp"
	"strings"
)

const (
	UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36"
)

type Parser struct {
	url    string
	result []byte
}

func NewParser(url string) Parser {
	return Parser{url: url}
}

func (p Parser) Parse() (*parser.ParseResult, error) {
	loc, err := utils.GetHeadersLocation(p.url)
	if err != nil {
		return nil, err
	}
	header := map[string]string{"User-Agent": UserAgent}
	//获取当前页面信息
	locurl := loc.URL.String()
	data, err := utils.SendRequest(locurl, header, nil)
	if err != nil {
		return nil, err
	}
	//获取页面script内容
	scriptContent, err := p.extractScriptContents(data)
	if err != nil {
		return nil, err
	}
	//解析script内容
	if p.result, err = p.parseScriptContent(scriptContent); err != nil {
		return nil, err
	}
	return p.parseResult()
}

func (p Parser) parseResult() (*parser.ParseResult, error) {
	var result Response
	if len(p.result) == 0 {
		return nil, errors.New("result is nil")
	}
	// 解析json数据
	err := json.Unmarshal(p.result, &result)
	if err != nil {
		return nil, err
	}
	return &parser.ParseResult{
		Data: &parser.Data{
			Author: result.Poster.Nick,
			Avatar: result.Poster.Avatar,
			Time:   utils.TimeStampToTime(result.Poster.Createtime, consts.TimeLayout),
			Title:  result.FeedDesc,
			Cover:  result.VideoCover,
			Url:    p.handleVideoUrl(result.VideoUrl),
		},
	}, nil
}

// 处理videoUrl
func (p Parser) handleVideoUrl(videoUrl string) string {
	if strings.Count(videoUrl, "v.weishi.qq.com") > 1 {
		return strings.Replace(videoUrl, "v.weishi.qq.com/v.weishi.qq.com", "v.weishi.qq.com", 1)
	}
	return videoUrl
}

func (p Parser) extractScriptContents(body []byte) (string, error) {
	if len(body) == 0 {
		return "", errors.New("body is nil")
	}
	scriptRegex := regexp.MustCompile(`<script>(.*?)</script>`)
	matches := scriptRegex.FindAllStringSubmatch(string(body), -1)

	for _, match := range matches {
		if len(match) < 2 || match[1] == "" {
			continue
		}
		if strings.Contains(match[1], "window.Vise.initState") {
			return strings.ReplaceAll(strings.ReplaceAll(match[1], "console.error('[Vise] fail to read initState.');", ""), "window.Vise.initState", "data"), nil
		}
	}
	return "", errors.New("script content not found")
}

// 解析script内容
func (p Parser) parseScriptContent(scriptContent string) ([]byte, error) {
	vm := goja.New()
	if _, err := vm.RunString(scriptContent); err != nil {
		return nil, err
	}
	// 获取视频信息
	result, err := vm.RunString("data.feedsList[0]")
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	if data, ok := result.Export().(map[string]interface{}); ok {
		jsonData, _ := json.Marshal(data)
		return jsonData, nil
	}
	return nil, errors.New("video data not found")
}
