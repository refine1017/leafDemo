package model

import (
	"gopkg.in/mgo.v2/bson"
	"server/base/msg"
)

type Player struct {
	Id            bson.ObjectId `bson:"_id" json:"id"`
	Username      string        `bson:"username" json:"username"`
	Nickname      string        `bson:"nickname" json:"nickname"`
	Level         uint32        `bson:"level" json:"level"`
	Vip           uint32        `bson:"vip" json:"vip"`
	LastLoginTime int64         `bson:"last_login_time" json:"last_login_time"`
}

func (p *Player) PlayerInfo() msg.PlayerInfo {
	return msg.PlayerInfo{
		Id:            p.Id.Hex(),
		Username:      p.Username,
		Nickname:      p.Nickname,
		Level:         p.Level,
		Vip:           p.Vip,
		LastLoginTime: p.LastLoginTime,
	}
}
