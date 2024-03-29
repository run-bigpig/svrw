package parser

type Data struct {
	Author string `json:"author"`
	Avatar string `json:"avatar"`
	Time   string `json:"time"`
	Title  string `json:"title"`
	Cover  string `json:"cover"`
	Url    string `json:"url"`
}

type ParseResult struct {
	Code int   `json:"code"`
	Data *Data `json:"data"`
}

type Parser interface {
	Parse() (*ParseResult, error)
}
