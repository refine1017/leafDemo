package login

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
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
	base.LoginRpc = skeleton.ChanRPCServer
}

func (m *Module) OnDestroy() {
	log.Release("Login Destroy")
}
