package parser

import (
	"encoding/json"
)

type Data struct {
	Author string `json:"author"`
	Avatar string `json:"avatar"`
	Time   string `json:"time"`
	Title  string `json:"title"`
	Cover  string `json:"cover"`
	Url    any    `json:"url"`
}

type ParseResult struct {
	*Data
}

type Parser interface {
	Parse() (*ParseResult, error)
}

func (pr *ParseResult) ToJson() []byte {
	jsonData, err := json.Marshal(pr)
	if err != nil {
		return nil
	}
	return jsonData
}
