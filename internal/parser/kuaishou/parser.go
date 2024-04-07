package kuaishou

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/run-bigpig/svrw/internal/parser"
	"github.com/run-bigpig/svrw/internal/utils"
	"time"
)

const (
	ApiUrl    = "https://v.m.chenzhongtech.com/rest/wd/photo/info"
	UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/"
	Cookie    = "did=%s; didv=%d"
)

type Parser struct {
	url    string
	result []byte
	header map[string]string
	req    *Request
}

func NewParser(url string) *Parser {
	return &Parser{url: url}
}

func (p *Parser) Parse() (*parser.ParseResult, error) {
	err := p.getQueryParams()
	if err != nil {
		return nil, err
	}
	p.result, err = utils.SendRequest(ApiUrl, p.header, p.req)
	if err != nil {
		return nil, err
	}
	return p.parseResult()
}

func (p *Parser) getQueryParams() error {
	loc, err := utils.GetHeadersLocation(p.url)
	if err != nil {
		return err
	}
	id := loc.URL.Query().Get("photoId")
	if id == "" {
		return errors.New("photoId not found")
	}
	did, err := p.getDid()
	if err != nil {
		return err
	}
	p.header = map[string]string{
		"User-Agent": UserAgent,
		"Referer":    loc.URL.String(),
		"Cookie":     fmt.Sprintf(Cookie, did, time.Now().UnixMilli()),
	}
	p.req = &Request{
		PhotoId:     id,
		IsLongVideo: false,
	}
	return nil
}

// 获取Did
func (p *Parser) getDid() (string, error) {
	cookies, err := utils.SendRequestGetCookie(p.url, map[string]string{"User-Agent": UserAgent}, nil)
	if err != nil {
		return "", errors.New("get did error")
	}
	for _, cookie := range cookies {
		if cookie.Name == "did" {
			return cookie.Value, nil
		}
	}
	return "", errors.New("did not found")
}

func (p *Parser) parseResult() (*parser.ParseResult, error) {
	var result Response
	err := json.Unmarshal(p.result, &result)
	if err != nil {
		return nil, err
	}
	if result.Result != 1 || len(result.Photos) == 0 {
		return nil, errors.New("parse result error")
	}
	photos := result.Photos[0]
	return &parser.ParseResult{
		Data: &parser.Data{
			Author: photos.UserName,
			Avatar: photos.HeadUrls[0].URL,
			Time:   utils.TimeStampToTime(photos.TimeStamp, "2006-01-02 15:04:05"),
			Title:  photos.Caption,
			Cover:  photos.CoverUrls[0].URL,
			Url:    photos.MainMvUrls[0].URL,
		},
	}, nil
}
