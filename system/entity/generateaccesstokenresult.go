package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GenerateAccessTokenResult struct{
    ExpiresIn	int	`json:"expires_in"`
    AccountId	string	`json:"account_id"`
    Country	string	`json:"country"`
    CountryUserInfo	[]GenerateAccessTokenCountryUserInfoResponseEntity	`json:"country_user_info"`
    AccountPlatform	string	`json:"account_platform"`
    AccessToken	string	`json:"access_token"`
    Account	string	`json:"account"`
    RefreshExpiresIn	string	`json:"refresh_expires_in"`
    RefreshToken	string	`json:"refresh_token"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GenerateAccessTokenDetailResponseEntity	`json:"detail"`
}
func (g GenerateAccessTokenResult) String() string {
    return lib.ObjectToString(g)
}
type GenerateAccessTokenDetailResponseEntity struct{
    
}