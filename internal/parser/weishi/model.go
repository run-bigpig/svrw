package weishi

type Poster struct {
	Avatar     string `json:"avatar"`
	Createtime int64  `json:"createtime"`
	Nick       string `json:"nick"`
}

type Response struct {
	FeedDesc     string `json:"feedDesc"`
	MaterialDesc string `json:"materialDesc"`
	Poster       Poster `json:"poster"`
	VideoCover   string `json:"videoCover"`
	VideoUrl     string `json:"videoUrl"`
}
