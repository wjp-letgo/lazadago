package etickets
import(
	"github.com/wjpxxx/letgo/lib"
	lazadaConfig "github.com/wjpxxx/lazadago/config"
	eticketsentity "github.com/wjpxxx/lazadago/etickets/entity"
)
//ETickets
type ETickets struct{
	Config *lazadaConfig.Config
}

//GetOrderItemsFromBarCode
//@Title E-Ticcket certificate query Open API
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=21&path=/eticket/code/query
func (s *ETickets) GetOrderItemsFromBarCode (code string) eticketsentity.GetOrderItemsFromBarCodeResult {
    method := "/eticket/code/query"
    params := lib.InRow{
      "code":code,
    }
    result := eticketsentity.GetOrderItemsFromBarCodeResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//RedeemOrderItems
//@Title Certificate Consume Open API
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=21&path=/eticket/code/consume
func (s *ETickets) RedeemOrderItems (bizType int,code string,outerId string,serialNum string,consumeNum int,storeId string,posId string) eticketsentity.RedeemOrderItemsResult {
    method := "/eticket/code/consume"
    params := lib.InRow{
      "biz_type":bizType,
      "code":code,
      "outer_id":outerId,
      "serial_num":serialNum,
      "consume_num":consumeNum,
    }
    if storeId!=""{
        params["store_id"]=storeId
    }
    if posId!=""{
        params["pos_id"]=posId
    }
    result := eticketsentity.RedeemOrderItemsResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//GlobalEticketMerchantMaAvailable
//@Title the callback interface before consume  code
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=21&path=/eticket/ma/available
func (s *ETickets) GlobalEticketMerchantMaAvailable (bizType int,code string,serialNum string,posId string,outerId string,consumeNum int,consumeStoreId string) eticketsentity.GlobalEticketMerchantMaAvailableResult {
    method := "/eticket/ma/available"
    params := lib.InRow{
      "biz_type":bizType,
      "code":code,
      "serial_num":serialNum,
      "outer_id":outerId,
      "consume_num":consumeNum,
      "consume_store_id":consumeStoreId,
    }
    if posId!=""{
        params["pos_id"]=posId
    }
    result := eticketsentity.GlobalEticketMerchantMaAvailableResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//GlobalEticketMerchantMaConsume
//@Title consume ma
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=21&path=/eticket/ma/consume
func (s *ETickets) GlobalEticketMerchantMaConsume (bizType int,serialNum string,posId string,outerId string,consumeNum int,code string,consumeStoreId string) eticketsentity.GlobalEticketMerchantMaConsumeResult {
    method := "/eticket/ma/consume"
    params := lib.InRow{
      "biz_type":bizType,
      "serial_num":serialNum,
      "outer_id":outerId,
      "consume_num":consumeNum,
      "code":code,
      "consume_store_id":consumeStoreId,
    }
    if posId!=""{
        params["pos_id"]=posId
    }
    result := eticketsentity.GlobalEticketMerchantMaConsumeResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//GlobalEticketMerchantMaFailsend
//@Title the callback interface when send code failed
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=21&path=/eticket/ma/failsend
func (s *ETickets) GlobalEticketMerchantMaFailsend (bizType int,subCode string,outerId string,subMsg string) eticketsentity.GlobalEticketMerchantMaFailsendResult {
    method := "/eticket/ma/failsend"
    params := lib.InRow{
      "biz_type":bizType,
      "sub_code":subCode,
      "outer_id":outerId,
      "sub_msg":subMsg,
    }
    result := eticketsentity.GlobalEticketMerchantMaFailsendResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//GlobalEticketMerchantMaQuery
//@Title the callback interface that query lazada platform ma
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=21&path=/eticket/ma/query
func (s *ETickets) GlobalEticketMerchantMaQuery (code string,sellerId int64,storeId int64) eticketsentity.GlobalEticketMerchantMaQueryResult {
    method := "/eticket/ma/query"
    params := lib.InRow{
      "code":code,
      "seller_id":sellerId,
    }
    if storeId>-1{
        params["store_id"]=storeId
    }
    result := eticketsentity.GlobalEticketMerchantMaQueryResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//GlobalEticketMerchantMaQueryTbMa
//@Title the callback interface that query tb ma
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=21&path=/eticket/ma/queryTbMa
func (s *ETickets) GlobalEticketMerchantMaQueryTbMa (code string) eticketsentity.GlobalEticketMerchantMaQueryTbMaResult {
    method := "/eticket/ma/queryTbMa"
    params := lib.InRow{
      "code":code,
    }
    result := eticketsentity.GlobalEticketMerchantMaQueryTbMaResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//GlobalEticketMerchantMaSend
//@Title the callback interface when merchant send code successful
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=21&path=/eticket/ma/send
func (s *ETickets) GlobalEticketMerchantMaSend (bizType int,isvMaList []eticketsentity.GlobalEticketMerchantMaSendIsvMaListRequestEntity,outerId string) eticketsentity.GlobalEticketMerchantMaSendResult {
    method := "/eticket/ma/send"
    params := lib.InRow{
      "biz_type":bizType,
      "isv_ma_list":isvMaList,
      "outer_id":outerId,
    }
    result := eticketsentity.GlobalEticketMerchantMaSendResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
