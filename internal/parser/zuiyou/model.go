package zuiyou

type AvatarUrls struct {
	Origin struct {
		Urls []string `json:"urls"`
	} `json:"origin"`
}

type Member struct {
	AvatarUrls AvatarUrls `json:"avatarUrls"`
	Name       string     `json:"name"`
	Sign       string     `json:"sign"`
}

type Videos struct {
	CoverUrls []string `json:"coverUrls"`
	URL       string   `json:"url"`
}

type Response struct {
	Content string            `json:"content"`
	Ct      int64             `json:"ct"`
	Member  Member            `json:"member"`
	Videos  map[string]Videos `json:"videos"`
}
