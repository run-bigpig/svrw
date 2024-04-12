package xiaohongshu

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dop251/goja"
	"github.com/run-bigpig/svrw/internal/consts"
	"github.com/run-bigpig/svrw/internal/parser"
	"github.com/run-bigpig/svrw/internal/utils"
	"log"
	"net/url"
	"regexp"
	"strings"
)

const (
	UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36"
	PlayUrl   = "https://sns-video-hw.xhscdn.com/%s"
)

type Parser struct {
	url    string
	noteId string
	result []byte
}

func NewParser(url string) *Parser {
	return &Parser{url: url}
}

func (p *Parser) Parse() (*parser.ParseResult, error) {
	err := p.checkUrl()
	if err != nil {
		return nil, err
	}
	err = p.getNoteId()
	if err != nil {
		return nil, err
	}
	//获取当前页面信息
	header := map[string]string{"User-Agent": UserAgent}
	data, err := utils.SendRequest(p.url, header, nil)
	if err != nil {
		return nil, err
	}
	//获取页面script内容
	scriptContent, err := p.extractScriptContents(data)
	if err != nil {
		return nil, err
	}
	//解析script内容
	p.result, err = p.parseScriptContent(scriptContent)
	if err != nil {
		return nil, err
	}
	return p.parseResult()
}

func (p *Parser) checkUrl() error {
	if strings.Contains(p.url, "xhslink.com") {
		loc, err := utils.GetHeadersLocation(p.url)
		if err != nil {
			return err
		}
		paths := strings.Split(strings.Trim(loc.URL.Path, "/"), "/")
		if len(paths) < 3 {
			return errors.New("url error")
		}
		p.url = fmt.Sprintf("https://www.xiaohongshu.com/explore/%s", paths[len(paths)-1])
	}
	return nil
}

func (p *Parser) getNoteId() error {
	u, err := url.Parse(p.url)
	if err != nil {
		return err
	}
	paths := strings.Split(strings.Trim(u.Path, "/"), "/")
	if len(paths) < 2 {
		return errors.New("noteId parser fail")
	}
	p.noteId = paths[len(paths)-1]
	return nil
}

func (p *Parser) extractScriptContents(body []byte) (string, error) {
	if len(body) == 0 {
		return "", errors.New("body is nil")
	}
	scriptRegex := regexp.MustCompile(`<script>(.*?)</script>`)
	matches := scriptRegex.FindAllStringSubmatch(string(body), -1)

	for _, match := range matches {
		if len(match) < 2 || match[1] == "" {
			continue
		}
		if strings.Contains(match[1], "window.__INITIAL_STATE__") {
			return strings.ReplaceAll(strings.ReplaceAll(match[1], "window.__INITIAL_STATE__", "var data"), "undefined", "1"), nil
		}
	}
	return "", errors.New("script content not found")
}

// 解析script内容
func (p *Parser) parseScriptContent(scriptContent string) ([]byte, error) {
	vm := goja.New()
	if _, err := vm.RunString(scriptContent); err != nil {
		log.Println(err.Error())
		return nil, err
	}
	// 获取视频信息
	runStr := fmt.Sprintf("data.note.noteDetailMap")
	result, err := vm.RunString(runStr)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	if data, ok := result.Export().(map[string]interface{}); ok {
		jsonData, _ := json.Marshal(data[p.noteId])
		return jsonData, nil
	}
	return nil, errors.New("video data not found")
}

func (p *Parser) parseResult() (*parser.ParseResult, error) {
	var result Response
	if err := json.Unmarshal(p.result, &result); err != nil {
		return nil, err
	}
	return &parser.ParseResult{
		Data: &parser.Data{
			Author: result.Note.User.Nickname,
			Avatar: result.Note.User.Avatar,
			Time:   utils.TimeStampToTime(result.Note.Time, consts.TimeLayout),
			Title:  result.Note.Title,
			Cover:  result.Note.ImageList[0].UrlDefault,
			Url:    fmt.Sprintf(PlayUrl, result.Note.Video.Consumer.OriginVideoKey),
		},
	}, nil
}
