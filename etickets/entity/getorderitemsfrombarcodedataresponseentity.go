package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetOrderItemsFromBarCodeDataResponseEntity struct{
    BizType	int	`json:"biz_type"`
    CertificateCode	string	`json:"certificate_code"`
    CodeStatus	string	`json:"code_status"`
    OuterId	string	`json:"outer_id"`
    StrartTime	int	`json:"strart_time"`
    EndTime	int	`json:"end_time"`
    TradeOrderId	int64	`json:"trade_order_id"`
    SerialNum	string	`json:"serial_num"`
    ItemList	[]GetOrderItemsFromBarCodeItemListResponseEntity	`json:"item_list"`
}
func (g GetOrderItemsFromBarCodeDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
