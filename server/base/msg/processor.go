package msg

import (
	"github.com/name5566/leaf/network/json"
)

var Processor = json.NewProcessor()

func init() {
	Processor.Register(&Hello{})
	Processor.Register(&DisconnectNotify{})
	Processor.Register(&RegisterReq{})
	Processor.Register(&RegisterRes{})
	Processor.Register(&LoginReq{})
	Processor.Register(&LoginRes{})
	Processor.Register(&TalkReq{})
	Processor.Register(&TalkRes{})
	Processor.Register(&TalkMsgNotify{})
}