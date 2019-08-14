package game

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"reflect"
	"server/base/model"
	"server/base/msg"
)

func init() {
	handler(&msg.Hello{}, handleHello)
	handler(&msg.TalkReq{}, handleTalk)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleHello(args []interface{}) {
	m := args[0].(*msg.Hello)
	a := args[1].(gate.Agent)

	log.Debug("hello %v", m.Name)

	a.WriteMsg(&msg.Hello{
		Name: "Client",
	})
}

func handleTalk(args []interface{}) {
	m := args[0].(*msg.TalkReq)
	a := args[1].(gate.Agent)

	if a.UserData() == nil {
		res := msg.TalkRes{}
		res.WithErrorString(msg.ErrPlayerNotLogin, "need login")
		a.WriteMsg(res)
		return
	}

	player := a.UserData().(*model.Player)

	notify := &msg.TalkMsgNotify{}
	notify.Msg = m.Msg

	agentMgr.Broadcast(notify)

	log.Debug("Player[%v] talk: %v", player.Id, m.Msg)
}