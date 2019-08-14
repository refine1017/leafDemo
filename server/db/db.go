package db

import (
	"github.com/name5566/leaf/db/mongodb"
	"server/base/conf"
)

var dbCtx *mongodb.DialContext

const (
	CollectionPlayers = "players"
)

func InitDbCxt() error {
	c, err := mongodb.Dial(conf.Server.MongoUrl, conf.Server.MongoSessionNum)
	if err != nil {
		return err
	}

	dbCtx = c

	return nil
}