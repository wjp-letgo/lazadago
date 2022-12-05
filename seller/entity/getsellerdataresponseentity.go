package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetSellerDataResponseEntity struct{
    NameCompany	string	`json:"name_company"`
    SellerId	int	`json:"seller_id"`
    Name	string	`json:"name"`
    ShortCode	string	`json:"short_code"`
    LogoUrl	string	`json:"logo_url"`
    Email	string	`json:"email"`
    Cb	bool	`json:"cb"`
    Location	string	`json:"location"`
    Status	string	`json:"status"`
    Verified	bool	`json:"verified"`
}
func (g GetSellerDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
