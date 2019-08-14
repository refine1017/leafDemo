package game

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"gopkg.in/mgo.v2/bson"
	"server/base/model"
)

var agentMgr = NewAgentManager()

type AgentManager struct {
	agents map[bson.ObjectId]gate.Agent
	index map[string]gate.Agent
}

func NewAgentManager() *AgentManager {
	return &AgentManager{
		agents: make(map[bson.ObjectId]gate.Agent),
		index: make(map[string]gate.Agent),
	}
}

func (am *AgentManager) AddAgent(agent gate.Agent, player *model.Player) {
	am.agents[player.Id] = agent
	am.index[player.Username] = agent

	log.Release("AddAgent Id=%v, Addr=%v", player.Id, agent.RemoteAddr())
}

func (am *AgentManager) DelAgent(agent gate.Agent) {
	data := agent.UserData()
	if data == nil {
		return
	}

	player := data.(*model.Player)

	delete(am.agents, player.Id)
	delete(am.index, player.Username)

	log.Release("DelAgent Id=%v, Addr=%v", player.Id, agent.RemoteAddr())
}

func (am *AgentManager) GetAgentByUsername(username string) gate.Agent {
	return am.index[username]
}

func (am *AgentManager) Broadcast(msg interface{}) {
	for _, agent := range am.agents {
		agent.WriteMsg(msg)
	}
}

func (am *AgentManager) Agents() map[bson.ObjectId]gate.Agent {
	return am.agents
}