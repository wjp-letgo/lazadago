package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type FreeShippingRegionsQueryDataResponseEntity struct{
    Name	string	`json:"name"`
    Value	string	`json:"value"`
}
func (g FreeShippingRegionsQueryDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
