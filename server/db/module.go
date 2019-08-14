package db

import (
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/module"
	"server/base"
	"server/base/conf"
)

var (
	skeleton = base.NewSkeleton()
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
	base.DBRpc = skeleton.ChanRPCServer

	if err := InitDbCxt(); err != nil {
		log.Fatal("InitDbCxt with err: %v", err)
	} else {
		log.Release("InitDbCtx %v with %v sessions", conf.Server.MongoUrl, conf.Server.MongoSessionNum)
	}
}

func (m *Module) OnDestroy() {
	if dbCtx != nil {
		dbCtx.Close()
		log.Release("DbCtx close")
	}
	log.Release("DB Destroy")
}