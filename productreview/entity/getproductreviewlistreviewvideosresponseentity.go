package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetProductReviewListReviewVideosResponseEntity struct{
    VideoUrl	string	`json:"video_url"`
    VideoCoverUrl	string	`json:"video_cover_url"`
}
func (g GetProductReviewListReviewVideosResponseEntity) String() string {
    return lib.ObjectToString(g)
}
