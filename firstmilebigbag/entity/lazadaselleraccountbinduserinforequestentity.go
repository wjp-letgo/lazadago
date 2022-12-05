package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type LazadaSellerAccountBindUserInfoRequestEntity struct{
    AppUserKey	string	`json:"appUserKey"`
}
func (g LazadaSellerAccountBindUserInfoRequestEntity) String() string {
    return lib.ObjectToString(g)
}
