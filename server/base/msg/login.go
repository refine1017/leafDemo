package msg

type PlayerInfo struct {
	Id            string `json:"id"`
	Nickname      string `json:"nickname"`
	Username      string `json:"username"`
	Level         uint32 `json:"level"`
	Vip           uint32 `json:"vip"`
	LastLoginTime int64  `json:"last_login_time"`
}

type LoginReq struct {
	Username string `json:"username"`
}

type LoginRes struct {
	Error
	PlayerInfo PlayerInfo `json:"player_info"`
}

type RegisterReq struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
}

type RegisterRes struct {
	Error
	PlayerInfo PlayerInfo `json:"player_info"`
}
