package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

type GetChannelcodeByFirstMileNoResult struct {
	Result    GetChannelcodeByFirstMileNoResultResponseEntity   `json:"result"`
	Type      string                                            `json:"type"`
	Code      string                                            `json:"code"`
	Message   string                                            `json:"message"`
	RequestId string                                            `json:"request_id"`
	Detail    []GetChannelcodeByFirstMileNoDetailResponseEntity `json:"detail"`
}

func (g GetChannelcodeByFirstMileNoResult) String() string {
	return lib.ObjectToString(g)
}

type GetChannelcodeByFirstMileNoDetailResponseEntity struct{}
