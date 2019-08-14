package gamedata

import (
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/recordfile"
	"reflect"
	"strings"
)

func readRf(st interface{}) *recordfile.RecordFile {
	rf, err := recordfile.New(st)
	if err != nil {
		log.Fatal("%v", err)
	}
	fn := strings.ToLower(reflect.TypeOf(st).Name() + ".txt")
	err = rf.Read("gamedata/" + fn)
	if err != nil {
		log.Fatal("%v: %v", fn, err)
	}

	return rf
}

func init() {
	log.Release("Read gamedata variables %v rows", Variables.NumRecord())
}
