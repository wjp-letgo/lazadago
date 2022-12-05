package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type RefreshAccessTokenResult struct{
    ExpiresIn	int	`json:"expires_in"`
    AccountId	string	`json:"account_id"`
    Country	string	`json:"country"`
    CountryUserInfoList	[]RefreshAccessTokenCountryUserInfoResponseEntity	`json:"country_user_info_list"`
    AccountPlatform	string	`json:"account_platform"`
    AccessToken	string	`json:"access_token"`
    Account	string	`json:"account"`
    RefreshExpiresIn	int	`json:"refresh_expires_in"`
    RefreshToken	string	`json:"refresh_token"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]RefreshAccessTokenDetailResponseEntity	`json:"detail"`
}
func (g RefreshAccessTokenResult) String() string {
    return lib.ObjectToString(g)
}

type RefreshAccessTokenDetailResponseEntity struct{
    
}