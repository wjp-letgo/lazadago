package productreview
import(
	"github.com/wjpxxx/letgo/lib"
	lazadaConfig "github.com/wjpxxx/lazadago/config"
	productreviewentity "github.com/wjpxxx/lazadago/productreview/entity"
)
//ProductReview
type ProductReview struct{
	Config *lazadaConfig.Config
}

//GetProductReviewList
//@Title get the review list for one seller
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=33&path=/review/seller/list
func (s *ProductReview) GetProductReviewList (itemId int64,orderId int64,startTime int,endTime int,contentFilter string,statusFilter string,pageSize int,current int) productreviewentity.GetProductReviewListResult {
    method := "/review/seller/list"
    params := lib.InRow{
      "item_id":itemId,
    }
    if orderId>-1{
        params["order_id"]=orderId
    }
    if startTime>-1{
        params["start_time"]=startTime
    }
    if endTime>-1{
        params["end_time"]=endTime
    }
    if contentFilter!=""{
        params["content_filter"]=contentFilter
    }
    if statusFilter!=""{
        params["status_filter"]=statusFilter
    }
    if pageSize>-1{
        params["page_size"]=pageSize
    }
    if current>-1{
        params["current"]=current
    }
    result := productreviewentity.GetProductReviewListResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//SubmitSellerReply
//@Title submit seller reply for customers review
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=33&path=/review/seller/reply/add
func (s *ProductReview) SubmitSellerReply (id int64,content string) productreviewentity.SubmitSellerReplyResult {
    method := "/review/seller/reply/add"
    params := lib.InRow{
      "id":id,
      "content":content,
    }
    result := productreviewentity.SubmitSellerReplyResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
