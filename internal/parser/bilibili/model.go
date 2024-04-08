package bilibili

type Owner struct {
	Mid  int    `json:"mid"`
	Name string `json:"name"`
	Face string `json:"face"`
}

type Page struct {
	Cid        int    `json:"cid"`
	Page       int    `json:"page"`
	Part       string `json:"part"`
	FirstFrame string `json:"first_frame"`
}

type Data struct {
	Bvid    string `json:"bvid"`
	Aid     int    `json:"aid"`
	Pic     string `json:"pic"`
	Title   string `json:"title"`
	Pubdate int64  `json:"pubdate"`
	Ctime   int64  `json:"ctime"`
	Desc    string `json:"desc"`
	Owner   Owner  `json:"owner"`
	Pages   []Page `json:"pages"`
}

type BaseResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type VideoData struct {
	Durl []PlayUrl `json:"durl"`
}

type PlayUrl struct {
	Page int    `json:"page,omitempty"`
	Url  string `json:"url"`
}

type PlayResponse struct {
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Data    VideoData `json:"data"`
}
