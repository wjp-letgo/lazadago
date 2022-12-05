package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetVideoResult struct{
    CoverUrl	string	`json:"cover_url"`
    VideoUrl	string	`json:"video_url"`
    Success	bool	`json:"success"`
    ResultCode	string	`json:"result_code"`
    State	string	`json:"state"`
    Title	string	`json:"title"`
    ResultMessage	string	`json:"result_message"`
    Type	string	`json:"type"`
    Code	string	`json:"code"`
    Message	string	`json:"message"`
    RequestId	string	`json:"request_id"`
    Detail	[]GetVideoDetailResponseEntity	`json:"detail"`
}
func (g GetVideoResult) String() string {
    return lib.ObjectToString(g)
}
type GetVideoDetailResponseEntity struct{}
