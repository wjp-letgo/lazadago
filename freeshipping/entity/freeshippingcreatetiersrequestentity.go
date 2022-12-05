package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type FreeShippingCreateTiersRequestEntity struct{
    Filter	string	`json:"filter"`
    Result	string	`json:"result"`
}
func (g FreeShippingCreateTiersRequestEntity) String() string {
    return lib.ObjectToString(g)
}
