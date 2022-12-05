package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetVideoQuotaResult struct{
    CapacitySize	int	`json:"capacity_size"`
    UsedSize	int	`json:"used_size"`
    Success	bool	`json:"success"`
    ResultCode	string	`json:"result_code"`
    ResultMessage	string	`json:"result_message"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GetVideoQuotaDetailResponseEntity	`json:"detail"`
}
func (g GetVideoQuotaResult) String() string {
    return lib.ObjectToString(g)
}
type GetVideoQuotaDetailResponseEntity struct{}
