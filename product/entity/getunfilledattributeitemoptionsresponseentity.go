package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetUnfilledAttributeItemOptionsResponseEntity struct {
	Name string `json:"name"`
}

func (g GetUnfilledAttributeItemOptionsResponseEntity) String() string {
	return lib.ObjectToString(g)
}
