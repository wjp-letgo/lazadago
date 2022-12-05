package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type ListFlexiComboDataListResponseEntity struct{
    OrderNumbers	int64	`json:"order_numbers"`
    PlatformChannel	string	`json:"platform_channel"`
    Name	string	`json:"name"`
    GiftSkus	[]ListFlexiComboGiftSkusResponseEntity	`json:"gift_skus"`
    StartTime	int	`json:"start_time"`
    OrderUsedNumbers	int	`json:"order_used_numbers"`
    DiscountType	string	`json:"discount_type"`
    EndTime	int	`json:"end_time"`
    Id	int	`json:"id"`
    DiscountValue	[]string	`json:"discount_value"`
    Status	string	`json:"status"`
    Apply	string	`json:"apply"`
    SampleSkus	[]ListFlexiComboSampleSkusResponseEntity	`json:"sample_skus"`
    CriteriaType	string	`json:"criteria_type"`
    Type	string	`json:"type"`
    CriteriaValue	[]string	`json:"criteria_value"`
    Stackable	bool	`json:"stackable"`
}
func (g ListFlexiComboDataListResponseEntity) String() string {
    return lib.ObjectToString(g)
}
