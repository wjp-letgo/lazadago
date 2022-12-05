package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type FreeShippingDeliveryOptionsQueryDataResponseEntity struct{
    Name	string	`json:"name"`
    Value	string	`json:"value"`
}
func (g FreeShippingDeliveryOptionsQueryDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
