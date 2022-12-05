package entity

import (
    "github.com/wjpxxx/letgo/lib"
)

type GetProductReviewListReviewsResponseEntity struct{
    Id	int	`json:"id"`
    ReviewType	string	`json:"review_type"`
    ReviewContent	string	`json:"review_content"`
    ReviewImages	[]string	`json:"review_images"`
    ReviewVideos	[]GetProductReviewListReviewVideosResponseEntity	`json:"review_videos"`
    SellerReply	string	`json:"seller_reply"`
}
func (g GetProductReviewListReviewsResponseEntity) String() string {
    return lib.ObjectToString(g)
}
