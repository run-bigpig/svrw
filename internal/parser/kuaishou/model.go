package kuaishou

type Request struct {
	PhotoId     string `json:"photoId"`
	IsLongVideo bool   `json:"isLongVideo"`
}

type MainMVUrls struct {
	CDN string `json:"cdn"`
	URL string `json:"url"`
}

type CoverUrls struct {
	CDN string `json:"cdn"`
	URL string `json:"url"`
}

type Photos struct {
	UserName      string       `json:"userName"`
	TimeStamp     int64        `json:"timestamp"`
	MainMvUrls    []MainMVUrls `json:"mainMvUrls"`
	CoverUrls     []CoverUrls  `json:"coverUrls"`
	WebpCoverUrls []CoverUrls  `json:"webpCoverUrls"`
	HeadUrls      []CoverUrls  `json:"headUrls"`
	Caption       string       `json:"caption"`
}

type Response struct {
	Result int      `json:"result"`
	Photos []Photos `json:"photos"`
}
