package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type LazadaBigbagUpdateUserInfoRequestEntity struct{
    AppUserKey	string	`json:"appUserKey"`
}
func (g LazadaBigbagUpdateUserInfoRequestEntity) String() string {
    return lib.ObjectToString(g)
}
