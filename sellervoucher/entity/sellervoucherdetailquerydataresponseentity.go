package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type SellerVoucherDetailQueryDataResponseEntity struct{
    CriteriaOverMoney	string	`json:"criteria_over_money"`
    Apply	string	`json:"apply"`
    VoucherType	string	`json:"voucher_type"`
    CollectStart	int	`json:"collect_start"`
    DisplayArea	string	`json:"display_area"`
    PeriodEndTime	int	`json:"period_end_time"`
    VoucherName	string	`json:"voucher_name"`
    VoucherDiscountType	string	`json:"voucher_discount_type"`
    OfferingMoneyValueOff	string	`json:"offering_money_value_off"`
    PeriodStartTime	int	`json:"period_start_time"`
    Limit	int	`json:"limit"`
    OrderUsedBudget	string	`json:"order_used_budget"`
    Currency	string	`json:"currency"`
    Id	int	`json:"id"`
    Issued	int	`json:"issued"`
    MaxDiscountOfferingMoneyValue	string	`json:"max_discount_offering_money_value"`
    VoucherCode	string	`json:"voucher_code"`
    OfferingPercentageDiscountOff	string	`json:"offering_percentage_discount_off"`
    Status	string	`json:"status"`
}
func (g SellerVoucherDetailQueryDataResponseEntity) String() string {
    return lib.ObjectToString(g)
}
