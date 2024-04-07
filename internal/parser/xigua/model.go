package xigua

type VideoResource struct {
	Dash Dash `json:"dash"`
}

type Dash struct {
	VideoList VideoList `json:"video_list"`
}

type VideoList struct {
	Video1 `json:"video_1"`
}

type Video1 struct {
	MainUrl string `json:"main_url"`
}

type UserInfo struct {
	AvatarURL string `json:"avatar_url"`
	Name      string `json:"name"`
}

type Video struct {
	PosterURL        string        `json:"poster_url"`
	Title            string        `json:"title"`
	UserDigg         int           `json:"user_digg"`
	UserInfo         UserInfo      `json:"user_info"`
	VideoResource    VideoResource `json:"videoResource"`
	VideoPublishTime string        `json:"video_publish_time"`
}

type PackerData struct {
	Video Video `json:"video"`
}

type Response struct {
	PackerData PackerData `json:"packerData"`
}
