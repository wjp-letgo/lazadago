package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetChannelcodeByFirstMileNoResultResponseEntity struct{
    Success	bool	`json:"success"`
    Module	[]GetChannelcodeByFirstMileNoModuleResponseEntity	`json:"module"`
    ErrorCode	string	`json:"errorCode"`
    ErrorMsg	string	`json:"errorMsg"`
}
func (g GetChannelcodeByFirstMileNoResultResponseEntity) String() string {
    return lib.ObjectToString(g)
}
type GetChannelcodeByFirstMileNoModuleResponseEntity struct{

}
func (g GetChannelcodeByFirstMileNoModuleResponseEntity) String() string {
    return lib.ObjectToString(g)
}