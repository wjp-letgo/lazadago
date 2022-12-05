package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetFulfillmentProductDetailSnSampleListResponseEntity struct{
    SampleSeq	string	`json:"sample_seq"`
    SampleDesc	string	`json:"sample_desc"`
    SampleRuleList	[]GetFulfillmentProductDetailSampleRuleListResponseEntity	`json:"sample_rule_list"`
}
func (g GetFulfillmentProductDetailSnSampleListResponseEntity) String() string {
    return lib.ObjectToString(g)
}
