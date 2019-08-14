package gamedata

import "strconv"

type Variable struct {
	Key string "index"
	Value string
}

var Variables = readRf(Variable{})

func InitPlayerLevel() uint32 {
	r := Variables.Index("INIT_PLAYER_LEVEL")
	if r == nil {
		return 0
	}

	v := r.(*Variable).Value

	level, err := strconv.Atoi(v)
	if err != nil {
		return 0
	}

	return uint32(level)
}

func InitPlayerVip() uint32 {
	r := Variables.Index("INIT_PLAYER_VIP")
	if r == nil {
		return 0
	}

	v := r.(*Variable).Value

	vip, err := strconv.Atoi(v)
	if err != nil {
		return 0
	}

	return uint32(vip)
}