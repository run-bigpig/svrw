package xiaohongshu

type User struct {
	Avatar   string `json:"avatar"`
	Nickname string `json:"nickname"`
	UserId   string `json:"userId"`
}

type Consumer struct {
	OriginVideoKey string `json:"originVideoKey"`
}

type Video struct {
	Consumer Consumer `json:"consumer"`
}

type Image struct {
	UrlDefault string `json:"urlDefault"`
}

type Note struct {
	NoteId     string  `json:"noteId"`
	Title      string  `json:"title"`
	Type       string  `json:"type"`
	IpLocation string  `json:"ipLocation"`
	Time       int64   `json:"time"`
	User       User    `json:"user"`
	ImageList  []Image `json:"imageList"`
	Video      Video   `json:"video"`
}

type Response struct {
	Note Note `json:"note"`
}
