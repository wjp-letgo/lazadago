package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type RefreshAccessTokenCountryUserInfoResponseEntity struct{
    Country	string	`json:"country"`
    SellerId	string	`json:"seller_id"`
    UserId	string	`json:"user_id"`
    ShortCode	string	`json:"short_code"`
}
func (g RefreshAccessTokenCountryUserInfoResponseEntity) String() string {
    return lib.ObjectToString(g)
}
