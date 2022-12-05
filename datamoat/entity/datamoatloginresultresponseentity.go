package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type DataMoatLoginResultResponseEntity struct {
	Msg     string `json:"msg"`
	Success bool   `json:"success"`
}

func (g DataMoatLoginResultResponseEntity) String() string {
	return lib.ObjectToString(g)
}
