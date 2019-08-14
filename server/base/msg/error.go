package msg

import "fmt"

type Error struct {
	Status uint32 `json:"status"`
	Error  string `json:"error"`
}

func (e *Error) WithError(status uint32, err error) {
	e.Status = status
	e.Error = fmt.Sprintf("%v", err)
}

func (e *Error) WithErrorString(status uint32, err string) {
	e.Status = status
	e.Error = err
}

// system error
const (
	ErrOk = iota
	ErrInternal
)

// common error
const (
	ErrPlayerNotLogin = iota + 100
)

// login error
const (
	ErrPlayerExist = iota + 200
)