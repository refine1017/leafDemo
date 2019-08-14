package main

import (
	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
	"server/base/conf"
	_ "server/base/gamedata"
	"server/db"
	"server/game"
	"server/gate"
	"server/login"
)

func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath

	leaf.Run(
		new(db.Module),
		new(game.Module),
		new(login.Module),
		new(gate.Module),
	)
}
