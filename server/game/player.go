package game

import (
	"github.com/name5566/leaf/gate"
	"server/base"
	"server/base/model"
)

func getPlayerFromAgent(agent gate.Agent) *model.Player {
	if agent.UserData() == nil {
		return nil
	}

	return agent.UserData().(*model.Player)
}

func savePlayer(player *model.Player) {
	base.DBRpc.Go("SavePlayerData", player)
}