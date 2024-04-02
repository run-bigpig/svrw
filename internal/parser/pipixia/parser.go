package pipixia

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/run-bigpig/svrw/internal/consts"
	"github.com/run-bigpig/svrw/internal/parser"
	"github.com/run-bigpig/svrw/internal/utils"
	"strings"
)

const ApiUrl = "https://is.snssdk.com/bds/cell/detail/?cell_type=1&aid=1319&app_name=super&cell_id=%s"

type Parser struct {
	url    string
	result []byte
}

func NewParser(url string) Parser {
	return Parser{url: url}
}

func (p Parser) Parse() (*parser.ParseResult, error) {
	itemId, err := p.getItemId()
	if err != nil {
		return nil, err
	}
	api := fmt.Sprintf(ApiUrl, itemId)
	p.result, err = utils.SendRequest(api, nil, nil)
	if err != nil {
		return nil, err
	}
	return p.parseResult()
}

// 解析url地址
func (p Parser) getItemId() (string, error) {
	loc, err := utils.GetHeadersLocation(p.url)
	if err != nil {
		return "", err
	}
	path := strings.Trim(loc.URL.Path, "/")
	hostSlice := strings.Split(path, "/")
	if len(hostSlice) < 2 {
		return "", errors.New("host not found")
	}
	return hostSlice[1], nil
}

// 解析结果
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
	// 解析结果
	if result.StatusCode != 0 {
		return nil, errors.New(result.Message)
	}
	return &parser.ParseResult{
		Data: &parser.Data{
			Author: result.Data.Data.Item.Author.Name,
			Avatar: result.Data.Data.Item.Author.Avatar.URLList[0].URL,
			Time:   utils.TimeStampToTime(result.Data.Data.Item.CreateTime, consts.TimeLayout),
			Title:  result.Data.Data.Item.Video.HashtagSchema[0].BaseHashtag.Intro,
			Cover:  result.Data.Data.Item.OriginVideoDownload.CoverImage.URLList[0].URL,
			Url:    result.Data.Data.Item.OriginVideoDownload.URLList[0].URL,
		},
	}, nil
}
