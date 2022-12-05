package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GlobalEticketMerchantMaQueryCertificateResponseEntity struct{
    LockedNum	int	`json:"locked_num"`
    BizType	int	`json:"biz_type"`
    CertificateCode	string	`json:"certificate_code"`
    InitialNum	int	`json:"initial_num"`
    AvailableNum	int	`json:"available_num"`
    ConsumeStatus	string	`json:"consume_status"`
    CodeStatus	string	`json:"code_status"`
    QrCodeUrl	string	`json:"qr_code_url"`
    OuterId	string	`json:"outer_id"`
    StartTime	int	`json:"start_time"`
    EndTime	int	`json:"end_time"`
    UsedNum	int	`json:"used_num"`
    Attributes	GlobalEticketMerchantMaQueryAttributesResponseEntity	`json:"attributes"`
}
func (g GlobalEticketMerchantMaQueryCertificateResponseEntity) String() string {
    return lib.ObjectToString(g)
}
type GlobalEticketMerchantMaQueryAttributesResponseEntity struct{}