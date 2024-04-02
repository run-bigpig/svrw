package douyin

type MigrationInfo struct {
	Ticket string `json:"ticket"`
	Source string `json:"source"`
}

type TTWidRequest struct {
	Region        string        `json:"region"`
	AID           int           `json:"aid"`
	NeedFID       bool          `json:"needFid"`
	Service       string        `json:"service"`
	MigrateInfo   MigrationInfo `json:"migrate_info"`
	CBUrlProtocol string        `json:"cbUrlProtocol"`
	Union         bool          `json:"union"`
}

type Author struct {
	AvatarThumb struct {
		URLList []string `json:"url_list"`
	} `json:"avatar_thumb"`
	Nickname string `json:"nickname"`
}

type Music struct {
	Author      string `json:"author"`
	AvatarLarge struct {
		Height  int      `json:"height"`
		URI     string   `json:"uri"`
		URLList []string `json:"url_list"`
		Width   int      `json:"width"`
	} `json:"avatar_large"`
	PlayURL struct {
		Height  int      `json:"height"`
		URI     string   `json:"uri"`
		URLKey  string   `json:"url_key"`
		URLList []string `json:"url_list"`
		Width   int      `json:"width"`
	} `json:"play_url"`
}

type OriginCover struct {
	Height  int      `json:"height"`
	URI     string   `json:"uri"`
	URLList []string `json:"url_list"`
	Width   int      `json:"width"`
}

type PlayAddr struct {
	DataSize int      `json:"data_size"`
	FileCS   string   `json:"file_cs"`
	FileHash string   `json:"file_hash"`
	Height   int      `json:"height"`
	URI      string   `json:"uri"`
	URLKey   string   `json:"url_key"`
	URLList  []string `json:"url_list"`
	Width    int      `json:"width"`
}

type Video struct {
	OriginCover OriginCover `json:"origin_cover"`
	PlayAddr    PlayAddr    `json:"play_addr"`
}

type AwemeDetail struct {
	Author     Author `json:"author"`
	CreateTime int64  `json:"create_time"`
	Desc       string `json:"desc"`
	Music      Music  `json:"music"`
	Video      Video  `json:"video"`
}

type Response struct {
	AwemeDetail AwemeDetail `json:"aweme_detail"`
	StatusCode  int         `json:"status_code"`
}
