package db

import (
	"github.com/name5566/leaf/log"
	"gopkg.in/mgo.v2/bson"
	"server/base/conf"
	"server/base/model"
	"server/base/gamedata"
)

func rpcFindPlayerByUsername(args []interface{}) []interface{} {
	username := args[0]

	r := dbCtx.Ref()
	defer dbCtx.UnRef(r)

	num, err := r.DB(conf.Server.MongoDB).C(CollectionPlayers).Find(bson.M{"username": username}).Count()
	if err != nil {
		return []interface{}{false, err}
	}

	if num == 0 {
		return []interface{}{false, nil}
	}

	return []interface{}{true, nil}
}

func rpcCreatePlayerData(args []interface{}) []interface{} {
	username, _ := args[0].(string)
	nickname, _ := args[1].(string)

	r := dbCtx.Ref()
	defer dbCtx.UnRef(r)

	var player = &model.Player{}
	player.Id = bson.NewObjectId()
	player.Username = username
	player.Nickname = nickname
	player.Level = gamedata.InitPlayerLevel()
	player.Vip = gamedata.InitPlayerVip()

	err := r.DB(conf.Server.MongoDB).C(CollectionPlayers).Insert(player)
	if err != nil {
		return []interface{}{nil, err}
	}

	return []interface{}{player, nil}
}

func rpcLoadPlayerDataByUsername(args []interface{}) []interface{} {
	username := args[0]

	r := dbCtx.Ref()
	defer dbCtx.UnRef(r)

	var player = &model.Player{}

	err := r.DB(conf.Server.MongoDB).C(CollectionPlayers).Find(bson.M{"username": username}).One(player)
	if err != nil {
		return []interface{}{nil, err}
	}

	return []interface{}{player, nil}
}

func rpcSavePlayerData(args []interface{}) {
	player := args[0].(*model.Player)

	r := dbCtx.Ref()
	defer dbCtx.UnRef(r)

	err := r.DB(conf.Server.MongoDB).C(CollectionPlayers).UpdateId(player.Id, player)
	if err != nil {
		log.Error("SavePlayer %v with err: %v", player.Id, err)
	} else {
		log.Release("SavePlayer %v success", player.Id)
	}
}