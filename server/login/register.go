package login

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"server/base"
	"server/base/model"
	"server/base/msg"
)

func handleRegister(args []interface{}) {
	m := args[0].(*msg.RegisterReq)
	a := args[1].(gate.Agent)

	res := &msg.RegisterRes{}

	var CreatePlayer = func() {
		skeleton.AsynCall(base.DBRpc, "CreatePlayerData", m.Username, m.Nickname, func(ret []interface{}, err error) {
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

			log.Release("Register player username=%v, nickname=%v, id=%v", m.Username, m.Nickname, player.Id)
		})
	}

	skeleton.AsynCall(base.DBRpc, "FindPlayerByUsername", m.Username, func(ret []interface{}, err error) {
		if err != nil {
			res.WithError(msg.ErrInternal, err)
			a.WriteMsg(res)
			return
		}

		result, _ := ret[0].(bool)
		err, _ = ret[1].(error)
		if err != nil {
			res.WithError(msg.ErrInternal, err)
			a.WriteMsg(res)
			return
		}

		if result {
			res.WithErrorString(msg.ErrPlayerExist, "Player is exists")
			a.WriteMsg(res)
			return
		}

		CreatePlayer()
	})
}
