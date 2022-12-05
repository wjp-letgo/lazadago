package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type FreeShippingListDataListResponseEntity struct{
    BudgetType	string	`json:"budget_type"`
    TemplateType	string	`json:"template_type"`
    UsedBudgetValue	string	`json:"used_budget_value"`
    Apply	string	`json:"apply"`
    PeriodEndTime	int	`json:"period_end_time"`
    TemplateCode	string	`json:"template_code"`
    CategoryName	string	`json:"category_name"`
    BudgetValue	string	`json:"budget_value"`
    PromotionName	string	`json:"promotion_name"`
    PeriodType	string	`json:"period_type"`
    RegionType	string	`json:"region_type"`
    PeriodStartTime	int	`json:"period_start_time"`
    PlatformChannel	string	`json:"platform_channel"`
    CampaignTag	string	`json:"campaign_tag"`
    RegionValue	[]string	`json:"region_value"`
    Currency	string	`json:"currency"`
    Id	int	`json:"id"`
    DeliveryOption	string	`json:"delivery_option"`
    PromoTier	FreeShippingListPromoTierResponseEntity	`json:"promo_tier"`
    Status	string	`json:"status"`
}
func (g FreeShippingListDataListResponseEntity) String() string {
    return lib.ObjectToString(g)
}
