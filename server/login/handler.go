package login

import (
	"reflect"
	"server/base/msg"
)

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handler(&msg.RegisterReq{}, handleRegister)
	handler(&msg.LoginReq{}, handleLogin)
}
