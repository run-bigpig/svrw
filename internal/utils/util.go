package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
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

// TimeStampToTime 时间戳转换为指定格式的时间
func TimeStampToTime(timeStamp int64, format string) string {
	t := time.Unix(timeStamp, 0)
	return t.Format(format)
}
