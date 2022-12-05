package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetNewBagNumberDataResponseEntity struct{
    LzdBagNumber	string	`json:"lzd_bag_number"`
}
func (g GetNewBagNumberDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
