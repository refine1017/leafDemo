package game

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"server/base/msg"
	"time"
)

func init() {
	skeleton.RegisterChanRPC("LoginAgent", rpcLoginAgent)
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
	skeleton.RegisterChanRPC("KickAgent", rpcKickAgent)
}

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	log.Release("NewAgent %v", a.RemoteAddr())
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	agentMgr.DelAgent(a)
}

func rpcLoginAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	player := getPlayerFromAgent(a)
	if player == nil {
		return
	}

	// if duplicate login, disconnect old agent
	if agent := agentMgr.GetAgentByUsername(player.Username); agent != nil {
		if agent.RemoteAddr() != a.RemoteAddr() {
			rpcKickAgent([]interface{}{agent, "duplicate login"})
		}
	}

	player.LastLoginTime = time.Now().Unix()
	savePlayer(player)

	agentMgr.AddAgent(a, player)
}

func rpcKickAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	r := args[1].(string)
	a.WriteMsg(&msg.DisconnectNotify{
		Reason: r,
	})
	a.Close()
	log.Release("Kick %v with reason %v", a.RemoteAddr(), r)
}
