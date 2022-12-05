package returnandrefund

import(
	"github.com/wjpxxx/letgo/lib"
	lazadaConfig "github.com/wjpxxx/lazadago/config"
	returnandrefundentity "github.com/wjpxxx/lazadago/returnandrefund/entity"
)
//ReturnAndRefund
type ReturnAndRefund struct{
	Config *lazadaConfig.Config
}

//GetReverseOrderDetail
//@Title Get the detailed information for a specific reverse order
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=26&path=/order/reverse/return/detail/list
func (s *ReturnAndRefund) GetReverseOrderDetail (reverseOrderId int64) returnandrefundentity.GetReverseOrderDetailResult {
    method := "/order/reverse/return/detail/list"
    params := lib.InRow{
      "reverse_order_id":reverseOrderId,
    }
    result := returnandrefundentity.GetReverseOrderDetailResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//GetReverseOrderHistoryList
//@Title Get the communication history of the reverse order
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=26&path=/order/reverse/return/history/list
func (s *ReturnAndRefund) GetReverseOrderHistoryList (reverseOrderLineId int64,pageSize int,pageNumber int) returnandrefundentity.GetReverseOrderHistoryListResult {
    method := "/order/reverse/return/history/list"
    params := lib.InRow{
      "reverse_order_line_id":reverseOrderLineId,
    }
    if pageSize>-1{
        params["page_size"]=pageSize
    }
    if pageNumber>-1{
        params["page_number"]=pageNumber
    }
    result := returnandrefundentity.GetReverseOrderHistoryListResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//GetReverseOrderReasonList
//@Title Get the list of reject reason. Need to be used in all refuse refund actions
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=26&path=/order/reverse/reason/list
func (s *ReturnAndRefund) GetReverseOrderReasonList (reverseOrderLineId int64) returnandrefundentity.GetReverseOrderReasonListResult {
    method := "/order/reverse/reason/list"
    params := lib.InRow{
      "reverse_order_line_id":reverseOrderLineId,
    }
    result := returnandrefundentity.GetReverseOrderReasonListResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//GetReverseOrdersForSeller
//@Title Use this API to get the list of items for a range of reverse orders.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=26&path=/reverse/getreverseordersforseller
func (s *ReturnAndRefund) GetReverseOrdersForSeller (ofcStatusList []string,reverseOrderId int64,tradeOrderId int64,pageSize int,reverseStatusList []string,pageNo int,returnToType string,disputeInProgress bool) returnandrefundentity.GetReverseOrdersForSellerResult {
    method := "/reverse/getreverseordersforseller"
    params := lib.InRow{
      "page_size":pageSize,
      "page_no":pageNo,
    }
    if reverseOrderId>-1{
        params["reverse_order_id"]=reverseOrderId
    }
    if tradeOrderId>-1{
        params["trade_order_id"]=tradeOrderId
    }
    if returnToType!=""{
        params["return_to_type"]=returnToType
    }
    result := returnandrefundentity.GetReverseOrdersForSellerResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//InitReverseOrderCancel
//@Title Seller initiates a cancelation
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=26&path=/order/reverse/cancel/create
func (s *ReturnAndRefund) InitReverseOrderCancel (orderItemIdList []string,orderId int64,reasonId string) returnandrefundentity.InitReverseOrderCancelResult {
    method := "/order/reverse/cancel/create"
    params := lib.InRow{
      "order_item_id_list":orderItemIdList,
      "order_id":orderId,
      "reason_id":reasonId,
    }
    result := returnandrefundentity.InitReverseOrderCancelResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//ReverseOrderCancelValidate
//@Title Seller can check whether the order can be canceled through this API and get corresponding reasons if not.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=26&path=/order/reverse/cancel/validate
func (s *ReturnAndRefund) ReverseOrderCancelValidate (orderId string,orderItemIdList []string) returnandrefundentity.ReverseOrderCancelValidateResult {
    method := "/order/reverse/cancel/validate"
    params := lib.InRow{
      "order_id":orderId,
      "order_item_id_list":orderItemIdList,
    }
    result := returnandrefundentity.ReverseOrderCancelValidateResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//ReverseOrderReturnUpdate
//@Title Seller can use this API to action on return and refund related.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=26&path=/order/reverse/return/update
func (s *ReturnAndRefund) ReverseOrderReturnUpdate (action string,reverseOrderId int64,reverseOrderItemIds []int64,reasonId int64,comment string,imageInfo []returnandrefundentity.ReverseOrderReturnUpdateImageInfoRequestEntity) returnandrefundentity.ReverseOrderReturnUpdateResult {
    method := "/order/reverse/return/update"
    params := lib.InRow{
      "action":action,
      "reverse_order_id":reverseOrderId,
      "reverse_order_item_ids":reverseOrderItemIds,
    }
    if reasonId>-1{
        params["reason_id"]=reasonId
    }
    if comment!=""{
        params["comment"]=comment
    }
    result := returnandrefundentity.ReverseOrderReturnUpdateResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
