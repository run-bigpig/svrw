package douyin

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
	"os"
	"strings"
)

const (
	ApiUrl    = "https://www.douyin.com/aweme/v1/web/aweme/detail/?aweme_id=%s&aid=1128&version_name=23.5.0&device_platform=android&os_version=2333"
	UserAgent = "Mozilla/5.0 (Linux; Android 10; SM-G960U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.181 Mobile Safari/537.36"
	RandStr   = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Referer   = "https://www.douyin.com"
	Cookie    = "msToken=%s;odin_tt=324fb4ea4a89c0c05827e18a1ed9cf9bf8a17f7705fcc793fec935b637867e2a5a9b8168c885554d029919117a18ba69; ttwid=%s; bd_ticket_guard_client_data=eyJiZC10aWNrZXQtZ3VhcmQtdmVyc2lvbiI6MiwiYmQtdGlja2V0LWd1YXJkLWNsaWVudC1jc3IiOiItLS0tLUJFR0lOIENFUlRJRklDQVRFIFJFUVVFU1QtLS0tLVxyXG5NSUlCRFRDQnRRSUJBREFuTVFzd0NRWURWUVFHRXdKRFRqRVlNQllHQTFVRUF3d1BZbVJmZEdsamEyVjBYMmQxXHJcbllYSmtNRmt3RXdZSEtvWkl6ajBDQVFZSUtvWkl6ajBEQVFjRFFnQUVKUDZzbjNLRlFBNUROSEcyK2F4bXAwNG5cclxud1hBSTZDU1IyZW1sVUE5QTZ4aGQzbVlPUlI4NVRLZ2tXd1FJSmp3Nyszdnc0Z2NNRG5iOTRoS3MvSjFJc3FBc1xyXG5NQ29HQ1NxR1NJYjNEUUVKRGpFZE1Cc3dHUVlEVlIwUkJCSXdFSUlPZDNkM0xtUnZkWGxwYmk1amIyMHdDZ1lJXHJcbktvWkl6ajBFQXdJRFJ3QXdSQUlnVmJkWTI0c0RYS0c0S2h3WlBmOHpxVDRBU0ROamNUb2FFRi9MQnd2QS8xSUNcclxuSURiVmZCUk1PQVB5cWJkcytld1QwSDZqdDg1czZZTVNVZEo5Z2dmOWlmeTBcclxuLS0tLS1FTkQgQ0VSVElGSUNBVEUgUkVRVUVTVC0tLS0tXHJcbiJ9"
)

type Parser struct {
	url    string
	header map[string]string
	result []byte
}

func NewParser(url string) *Parser {
	return &Parser{url: url, header: make(map[string]string)}
}

func (p *Parser) Parse() (*parser.ParseResult, error) {
	itemId, err := p.getItemId()
	if err != nil {
		return nil, err
	}
	xbogusUrl := fmt.Sprintf(ApiUrl, itemId)
	xbogus, err := p.sign(xbogusUrl, UserAgent)
	if err != nil {
		return nil, err
	}
	api := fmt.Sprintf("%s&X-Bogus=%s", xbogusUrl, xbogus)
	p.getQueryParams()
	p.result, err = utils.SendRequest(api, p.header, nil)
	if err != nil {
		return nil, err
	}
	return p.parseResult()
}

func (p *Parser) getItemId() (string, error) {
	loc, err := utils.GetHeadersLocation(p.url)
	if err != nil {
		return "", err
	}
	path := strings.Trim(loc.URL.Path, "/")
	slice := strings.Split(path, "/")
	if len(slice) < 2 {
		return "", errors.New("host not found")
	}
	return slice[1], nil
}

func (p *Parser) sign(urlApi, userAgent string) (string, error) {
	u, err := url.Parse(urlApi)
	if err != nil {
		return "", err
	}
	vm := goja.New()
	if _, err := vm.RunString(Xbogus); err != nil {
		return "", err
	}

	result, err := vm.RunString(fmt.Sprintf("sign('%s', '%s')", u.RawQuery, userAgent))
	if err != nil {
		return "", err
	}

	xbogus, ok := result.Export().(string)
	if !ok {
		return "", fmt.Errorf("Failed to convert result to string")
	}
	return xbogus, nil
}

func (p *Parser) openJSFile(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		return ""
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return ""
	}

	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return ""
	}
	return string(bs)
}

// 组装请求参数
func (p *Parser) getQueryParams() {
	msToken := utils.SubStr(utils.StrShuffle(RandStr), 0, 107)
	ttwid, err := p.getttwid()
	if err != nil {
		log.Println(err)
	}
	p.header["User-Agent"] = UserAgent
	p.header["Referer"] = Referer
	p.header["Cookie"] = fmt.Sprintf(Cookie, msToken, ttwid)
}

// 解析结果
func (p *Parser) parseResult() (*parser.ParseResult, error) {
	var data Response
	if len(p.result) == 0 {
		return nil, errors.New("result is nil")
	}
	err := json.Unmarshal(p.result, &data)
	if err != nil {
		return nil, err
	}
	if data.StatusCode != 0 {
		return nil, errors.New("parse result error")
	}
	return &parser.ParseResult{
		Data: &parser.Data{
			Author: data.AwemeDetail.Author.Nickname,
			Avatar: data.AwemeDetail.Author.AvatarThumb.URLList[0],
			Time:   utils.TimeStampToTime(data.AwemeDetail.CreateTime, consts.TimeLayout),
			Title:  data.AwemeDetail.Desc,
			Cover:  data.AwemeDetail.Video.OriginCover.URLList[0],
			Url:    data.AwemeDetail.Video.PlayAddr.URLList[0],
		},
	}, nil
}

// 获取ttwid
func (p *Parser) getttwid() (string, error) {
	ttwidurl := "https://ttwid.bytedance.com/ttwid/union/register/"
	data := TTWidRequest{
		Region:  "cn",
		AID:     1768,
		NeedFID: false,
		Service: "www.ixigua.com",
		MigrateInfo: MigrationInfo{
			Ticket: "",
			Source: "node",
		},
		CBUrlProtocol: "https",
		Union:         true,
	}
	header := map[string]string{
		"User-Agent":   UserAgent,
		"Content-Type": "application/json",
	}
	result, err := utils.SendRequestGetCookie(ttwidurl, header, data)
	if err != nil {
		return "", err
	}
	for _, v := range result {
		if v.Name == "ttwid" {
			return v.Value, nil
		}
	}
	return "", errors.New("ttwid not found")
}
