package msg

type Hello struct {
	Name string `json:"name"`
}

type DisconnectNotify struct {
	Reason string `json:"reason"`
}

type TalkReq struct {
	Msg string `json:"msg"`
}

type TalkRes struct {
	Error
}

type TalkMsgNotify struct {
	Msg string `json:"msg"`
}