package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type SetStatusToPackedByMarketplaceResult struct{
    Data	SetStatusToPackedByMarketplaceDataResponseEntity	`json:"data"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]SetStatusToPackedByMarketplaceDetailResponseEntity	`json:"detail"`
}
func (g SetStatusToPackedByMarketplaceResult) String() string {
    return lib.ObjectToString(g)
}

type SetStatusToPackedByMarketplaceDetailResponseEntity struct{
    
}
