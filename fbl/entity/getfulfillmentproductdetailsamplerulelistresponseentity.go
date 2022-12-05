package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetFulfillmentProductDetailSampleRuleListResponseEntity struct{
    RuleRegularExpression	string	`json:"rule_regular_expression"`
    RuleDesc	string	`json:"rule_desc"`
    RuleImgUrl	string	`json:"rule_img_url"`
    RuleSample	string	`json:"rule_sample"`
}
func (g GetFulfillmentProductDetailSampleRuleListResponseEntity) String() string {
    return lib.ObjectToString(g)
}
