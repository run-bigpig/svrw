package xigua

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
	Cookie    = "csrf_session_id=3857bf159ebd052f4961b709db2ff431; ttwid=1%7Cn-EVYT7ML38_Ikp84SXfSFXnE3JeIKrdMWcvde-Jf2o%7C1712475048%7Cb9ef20cbdb97a9a1ca427f2625cf434c39a6d4e41eccddc4bdd7153f8652c94e; support_webp=true; support_avif=true; msToken=79rvpqPEJIbsKe0kvY3unJVzvpLMLZYv9I9dF2uaIsZ31mNtu6OlVV6vMfsE8IFTaeDAqyKTz76y7VjWdqFklsahg_UlJSVUYFT8SBbM; ixigua-a-s=1; msToken=54YlH3Hr2CGTQ0VEiRl3oriQXP7bxUNVuwuxHfNJXg9gSYsqRUzliQEx-WGlXA5YCTXlUDEtTs_fYbKM5XkUzMgSx2iDMWugVziC6YEDIqtcV4VJ8Bza"
)

type Parser struct {
	url    string
	result []byte
}

func NewParser(url string) *Parser {
	return &Parser{url: url}
}

func (p *Parser) Parse() (*parser.ParseResult, error) {
	//获取当前页面信息
	header := map[string]string{"User-Agent": UserAgent, "Cookie": Cookie}
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

func (p *Parser) extractScriptContents(body []byte) (string, error) {
	if len(body) == 0 {
		return "", errors.New("body is nil")
	}
	scriptRegex := regexp.MustCompile(`<script id="SSR_HYDRATED_DATA" nonce="[^"]*">(.*?)</script>`)
	matches := scriptRegex.FindAllStringSubmatch(string(body), -1)

	for _, match := range matches {
		if len(match) < 2 || match[1] == "" {
			continue
		}
		if strings.Contains(match[1], "window._SSR_HYDRATED_DATA") {
			return strings.ReplaceAll(match[1], "window._SSR_HYDRATED_DATA", "data"), nil
		}
	}
	return "", errors.New("script content not found")
}

// 解析script内容
func (p *Parser) parseScriptContent(scriptContent string) ([]byte, error) {
	vm := goja.New()
	if _, err := vm.RunString(scriptContent); err != nil {
		return nil, err
	}
	// 获取视频信息
	result, err := vm.RunString("data.anyVideo.gidInformation")
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

func (p *Parser) parseResult() (*parser.ParseResult, error) {
	var result Response
	if err := json.Unmarshal(p.result, &result); err != nil {
		return nil, err
	}
	return &parser.ParseResult{
		Data: &parser.Data{
			Author: result.PackerData.Video.UserInfo.Name,
			Avatar: result.PackerData.Video.UserInfo.AvatarURL,
			Time:   utils.TimeStampToTime(utils.StringToInt64(result.PackerData.Video.VideoPublishTime), consts.TimeLayout),
			Title:  result.PackerData.Video.Title,
			Cover:  result.PackerData.Video.PosterURL,
			Url:    utils.Base64Decode(result.PackerData.Video.VideoResource.Normal.VideoList.Video1.MainUrl),
		},
	}, nil
}
