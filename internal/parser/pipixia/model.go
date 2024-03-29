package pipixia

type Response struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Prompt     string `json:"prompt"`
	Time       int64  `json:"time"`
	Data       Data   `json:"data"`
}

type Data struct {
	Data struct {
		Item Item `json:"item"`
	} `json:"data"`
}

type Item struct {
	CreateTime          int64  `json:"create_time"`
	Video               Video  `json:"video"`
	Author              Author `json:"author"`
	OriginVideoDownload Origin `json:"origin_video_download"`
}

type Video struct {
	HashtagSchema []Hashtag `json:"hashtag_schema"`
}

type Hashtag struct {
	BaseHashtag BaseHashtag `json:"base_hashtag"`
}

type BaseHashtag struct {
	Name  string `json:"name"`
	Intro string `json:"intro"`
}

type Author struct {
	Name   string `json:"name"`
	Avatar Avatar `json:"avatar"`
}

type Avatar struct {
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	URI          string `json:"uri"`
	URLList      []URL  `json:"url_list"`
	IsGIF        bool   `json:"is_gif"`
	DownloadList []URL  `json:"download_list"`
}

type URL struct {
	URL string `json:"url"`
}

type Origin struct {
	URI        string `json:"uri"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
	URLList    []URL  `json:"url_list"`
	CoverImage Cover  `json:"cover_image"`
}

type Cover struct {
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	URI          string `json:"uri"`
	URLList      []URL  `json:"url_list"`
	IsGIF        bool   `json:"is_gif"`
	DownloadList []URL  `json:"download_list"`
}
