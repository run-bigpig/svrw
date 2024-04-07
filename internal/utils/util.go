package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func SendRequest(url string, headers map[string]string, data interface{}) ([]byte, error) {
	var (
		req *http.Request
		err error
	)
	client := &http.Client{
		Timeout: time.Second * 50,
	}
	if data != nil {
		body, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		req, err = http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/json")
	} else {
		req, err = http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return nil, err
		}
	}
	// 设置请求头（如果有）
	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SendRequestGetCookie 发送请求并获取cookie
func SendRequestGetCookie(url string, headers map[string]string, data interface{}) ([]*http.Cookie, error) {
	var (
		req *http.Request
		err error
	)
	client := &http.Client{
		Timeout: time.Second * 50,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return nil
		},
	}
	if data != nil {
		body, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		req, err = http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/json")
	} else {
		req, err = http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return nil, err
		}
	}
	// 设置请求头（如果有）
	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return resp.Cookies(), nil
}

// TimeStampToTime 时间戳转换为指定格式的时间
func TimeStampToTime(timeStamp int64, format string) string {
	//timeStamp为毫秒级时间戳处理
	if timeStamp < 9999999999 {
		timeStamp = timeStamp * 1000
	}
	tm := time.Unix(timeStamp/1e3, 0)
	return tm.Format(format)
}

// GetHeadersLocation 获取重定向地址
func GetHeadersLocation(url string) (*http.Request, error) {
	resp, err := http.Head(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp.Request, nil
}

func StrShuffle(s string) string {
	runes := []rune(s)
	for i := len(runes) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func SubStr(s string, start, length int) string {
	if start < 0 {
		start = 0
	}
	if start >= len(s) {
		return ""
	}
	end := start + length
	if end > len(s) {
		end = len(s)
	}
	return s[start:end]
}

// Base64Decode base64解码
func Base64Decode(s string) string {
	de, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return ""
	}
	return string(de)
}

func StringToInt64(s string) int64 {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return int64(i)
}
