package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetFlexiComboDetailsDataResponseEntity struct{
    OrderUsedNumbers	int	`json:"order_used_numbers"`
    Apply	string	`json:"apply"`
    SampleSkus	[]GetFlexiComboDetailsSampleSkusResponseEntity	`json:"sample_skus"`
    CriteriaType	string	`json:"criteria_type"`
    Type	string	`json:"type"`
    CriteriaValue	[]string	`json:"criteria_value"`
    OrderNumbers	int64	`json:"order_numbers"`
    PlatformChannel	string	`json:"platform_channel"`
    Name	string	`json:"name"`
    GiftSkus	[]GetFlexiComboDetailsGiftSkusResponseEntity	`json:"gift_skus"`
    DiscountType	string	`json:"discount_type"`
    StartTime	int	`json:"start_time"`
    EndTime	int	`json:"end_time"`
    Id	int	`json:"id"`
    DiscountValue	[]string	`json:"discount_value"`
    Status	string	`json:"status"`
    Stackable	bool	`json:"stackable"`
}
func (g GetFlexiComboDetailsDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
