package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetOrderTraceLogisticDetailInfoListResponseEntity struct{
    PackageLocationName	string	`json:"package_location_name"`
    EventDate	string	`json:"event_date"`
    DetailType	string	`json:"detail_type"`
    ProofImages	[]GetOrderTraceProofImagesResponseEntity	`json:"proof_images"`
    ReceiveTime	int	`json:"receive_time"`
    StatusCode	string	`json:"status_code"`
    Icon	string	`json:"icon"`
    EventTime	int	`json:"event_time"`
    Description	string	`json:"description"`
    Title	string	`json:"title"`
}
func (g GetOrderTraceLogisticDetailInfoListResponseEntity) String() string {
    return lib.ObjectToString(g)
}

type GetOrderTraceProofImagesResponseEntity struct{
    
}