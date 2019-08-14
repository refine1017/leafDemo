package game

import (
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/module"
	"server/base"
)

var (
	skeleton = base.NewSkeleton()
)

type Module struct {
	*module.Skeleton
	agentMgr *AgentManager
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
	base.GameRpc = skeleton.ChanRPCServer
}

func (m *Module) OnDestroy() {
	log.Release("Game Destroy")
}
