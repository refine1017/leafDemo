package login

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"server/base"
	"server/base/model"
	"server/base/msg"
)

func handleLogin(args []interface{}) {
	m := args[0].(*msg.LoginReq)
	a := args[1].(gate.Agent)

	res := &msg.LoginRes{}

	// if player has login
	if a.UserData() != nil {
		player := a.UserData().(*model.Player)

		// if login different account, will close agent
		if player.Username != m.Username {
			base.GameRpc.Go("KickAgent", a, "login different account")
			return
		}

		res.PlayerInfo = player.PlayerInfo()
		a.WriteMsg(res)

		base.GameRpc.Go("LoginAgent", a)

		log.Release("ReLogin player username=%v, nickname=%v, id=%v", m.Username, player.Nickname, player.Id)
		return
	}

	// load player data
	skeleton.AsynCall(base.DBRpc, "LoadPlayerDataByUsername", m.Username, func(ret []interface{}, err error) {
		if err != nil {
			res.WithError(msg.ErrInternal, err)
			a.WriteMsg(res)
			return
		}

		player, _ := ret[0].(*model.Player)
		err, _ = ret[1].(error)
		if err != nil {
			res.WithError(msg.ErrInternal, err)
			a.WriteMsg(res)
			return
		}

		res.PlayerInfo = player.PlayerInfo()
		a.WriteMsg(res)

		a.SetUserData(player)

		base.GameRpc.Go("LoginAgent", a)

		log.Release("Login player username=%v, nickname=%v, id=%v", m.Username, player.Nickname, player.Id)
	})
}
