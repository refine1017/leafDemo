package gate

import (
	"server/base"
	"server/base/msg"
)

func initRouter() {
	// login
	msg.Processor.SetRouter(&msg.RegisterReq{}, base.LoginRpc)
	msg.Processor.SetRouter(&msg.LoginReq{}, base.LoginRpc)

	// game
	msg.Processor.SetRouter(&msg.Hello{}, base.GameRpc)
	msg.Processor.SetRouter(&msg.TalkReq{}, base.GameRpc)
}
