package order

import(
	"github.com/wjpxxx/letgo/lib"
	lazadaConfig "github.com/wjpxxx/lazadago/config"
	orderentity "github.com/wjpxxx/lazadago/order/entity"
    "strings"
)
//Order
type Order struct{
	Config *lazadaConfig.Config
}

//GetAwbDocumentHtml
//@Title Use this API to retrieve order-related documents, only for shipping labels.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.bd786bbeNJUDaJ#/api?cid=8&path=/order/document/awb/html/get
func (s *Order) GetAwbDocumentHtml (orderItemIds string) orderentity.GetAwbDocumentHtmlResult {
    method := "/order/document/awb/html/get"
    params := lib.InRow{
      "order_item_ids":orderItemIds,
    }
    result := orderentity.GetAwbDocumentHtmlResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}

//GetAwbDocumentPDF
//@Title Use this API to retrieve order-related documents, only for shipping labels.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.bd786bbeNJUDaJ#/api?cid=8&path=/order/document/awb/pdf/get
func (s *Order) GetAwbDocumentPDF (orderItemIds string) orderentity.GetAwbDocumentPDFResult {
    method := "/order/document/awb/pdf/get"
    params := lib.InRow{
      "order_item_ids":orderItemIds,
    }
    result := orderentity.GetAwbDocumentPDFResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//GetDocument
//@Title Use this API to retrieve order-related documents, including invoices and shipping labels.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.bd786bbeNJUDaJ#/api?cid=8&path=/order/document/get
func (s *Order) GetDocument (docType string,orderItemIds string) orderentity.GetDocumentResult {
    method := "/order/document/get"
    params := lib.InRow{
      "doc_type":docType,
      "order_item_ids":orderItemIds,
    }
    result := orderentity.GetDocumentResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//GetFailureReasons
//@Title Use this API to get additional error context for the SetStatusToCanceled API.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.bd786bbeNJUDaJ#/api?cid=8&path=/order/failure_reason/get
func (s *Order) GetFailureReasons () orderentity.GetFailureReasonsResult {
    method := "/order/failure_reason/get"
    params := lib.InRow{}
    result := orderentity.GetFailureReasonsResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//GetMultipleOrderItems
//@Title Use this API to get the item information of one or more orders.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.bd786bbeNJUDaJ#/api?cid=8&path=/orders/items/get
func (s *Order) GetMultipleOrderItems (orderIds []int64) orderentity.GetMultipleOrderItemsResult {
    ids:=lib.Int64ArrayToArrayString(orderIds)
    method := "/orders/items/get"
    params := lib.InRow{
      "order_ids":"["+strings.Join(ids,",")+"]",
    }
    result := orderentity.GetMultipleOrderItemsResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//GetOVOOrders
//@Title OVO order query
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.bd786bbeNJUDaJ#/api?cid=8&path=/orders/ovo/get
func (s *Order) GetOVOOrders (tradeOrderIds string) orderentity.GetOVOOrdersResult {
    method := "/orders/ovo/get"
    params := lib.InRow{
      "tradeOrderIds":tradeOrderIds,
    }
    result := orderentity.GetOVOOrdersResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//GetOrder
//@Title Use this API to get the list of items for a single order.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.bd786bbeNJUDaJ#/api?cid=8&path=/order/get
func (s *Order) GetOrder (orderId int64) orderentity.GetOrderResult {
    method := "/order/get"
    params := lib.InRow{
      "order_id":orderId,
    }
    result := orderentity.GetOrderResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//GetOrderItems
//@Title Use this API to get the item information of an order.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.bd786bbeNJUDaJ#/api?cid=8&path=/order/items/get
func (s *Order) GetOrderItems (orderId int64) orderentity.GetOrderItemsResult {
    method := "/order/items/get"
    params := lib.InRow{
      "order_id":orderId,
    }
    result := orderentity.GetOrderItemsResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}

//GetOrders
//@Title Use this API to get the list of items for a range of orders1..
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.bd786bbeNJUDaJ#/api?cid=8&path=/orders/get
func (s *Order) GetOrdersByUpdateTime (start,end int,offset int,limit int, status string) orderentity.GetOrdersResult {
    updateBeforeStr:=lib.TimeToFormat(end,"2006-01-02T15:04:05+0800")
    updateAfterStr:=lib.TimeToFormat(start,"2006-01-02T15:04:05+0800")
    return s.GetOrders(updateBeforeStr,"",offset,limit,updateAfterStr,"","","",status)
}

//GetOrders
//@Title Use this API to get the list of items for a range of orders1..
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.bd786bbeNJUDaJ#/api?cid=8&path=/orders/get
func (s *Order) GetOrders (updateBefore string,sortDirection string,offset int,limit int,updateAfter string,sortBy string,createdBefore string,createdAfter string,status string) orderentity.GetOrdersResult {
    method := "/orders/get"
    params := lib.InRow{
    }
    if updateBefore!=""{
        params["update_before"]=updateBefore
    }
    if sortDirection!=""{
        params["sort_direction"]=sortDirection
    }
    if offset>-1{
        params["offset"]=offset
    }
    if limit>-1{
        params["limit"]=limit
    }
    if updateAfter!=""{
        params["update_after"]=updateAfter
    }
    if sortBy!=""{
        params["sort_by"]=sortBy
    }
    if createdBefore!=""{
        params["created_before"]=createdBefore
    }
    if createdAfter!=""{
        params["created_after"]=createdAfter
    }
    if status!=""{
        params["status"]=status
    }
    result := orderentity.GetOrdersResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//SetInvoiceNumber
//@Title Use this API to set the invoice number for the specified order.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.bd786bbeNJUDaJ#/api?cid=8&path=/order/invoice_number/set
func (s *Order) SetInvoiceNumber (orderItemId int64,invoiceNumber string) orderentity.SetInvoiceNumberResult {
    method := "/order/invoice_number/set"
    params := lib.InRow{
      "order_item_id":orderItemId,
      "invoice_number":invoiceNumber,
    }
    result := orderentity.SetInvoiceNumberResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//SetRepack
//@Title Use this API to mark a package item as being repack.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.bd786bbeNJUDaJ#/api?cid=8&path=/order/repack
func (s *Order) SetRepack (packageId string) orderentity.SetRepackResult {
    method := "/order/repack"
    params := lib.InRow{
      "package_id":packageId,
    }
    result := orderentity.SetRepackResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//SetStatusToCanceled
//@Title Use this API to cancel a single order item.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.bd786bbeNJUDaJ#/api?cid=8&path=/order/cancel
func (s *Order) SetStatusToCanceled (reasonDetail string,reasonId int64,orderItemId int64) orderentity.SetStatusToCanceledResult {
    method := "/order/cancel"
    params := lib.InRow{
      "reason_id":reasonId,
      "order_item_id":orderItemId,
    }
    if reasonDetail!=""{
        params["reason_detail"]=reasonDetail
    }
    result := orderentity.SetStatusToCanceledResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//SetStatusToPackedByMarketplace
//@Title Use this API to mark an order item as being packed.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.bd786bbeNJUDaJ#/api?cid=8&path=/order/pack
func (s *Order) SetStatusToPackedByMarketplace (shippingProvider string,deliveryType string,orderItemIds string) orderentity.SetStatusToPackedByMarketplaceResult {
    method := "/order/pack"
    params := lib.InRow{
      "shipping_provider":shippingProvider,
      "delivery_type":deliveryType,
      "order_item_ids":orderItemIds,
    }
    result := orderentity.SetStatusToPackedByMarketplaceResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//SetStatusToReadyToShip
//@Title Use this API to mark an order item as being ready to ship.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.bd786bbeNJUDaJ#/api?cid=8&path=/order/rts
func (s *Order) SetStatusToReadyToShip (deliveryType string,orderItemIds string,shipmentProvider string,trackingNumber string) orderentity.SetStatusToReadyToShipResult {
    method := "/order/rts"
    params := lib.InRow{
      "delivery_type":deliveryType,
      "order_item_ids":orderItemIds,
      "shipment_provider":shipmentProvider,
      "tracking_number":trackingNumber,
    }
    result := orderentity.SetStatusToReadyToShipResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//SetStatusToSOFDelivered
//@Title Use this API to mark an sof order item as being delivered.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.bd786bbeNJUDaJ#/api?cid=8&path=/order/sof/delivered
func (s *Order) SetStatusToSOFDelivered (orderItemIds string) orderentity.SetStatusToSOFDeliveredResult {
    method := "/order/sof/delivered"
    params := lib.InRow{
      "order_item_ids":orderItemIds,
    }
    result := orderentity.SetStatusToSOFDeliveredResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//SetStatusToSOFFailedDelivery
//@Title Use this API to mark an sof order item as being delivered failed
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.bd786bbeNJUDaJ#/api?cid=8&path=/order/sof/failed_delivery
func (s *Order) SetStatusToSOFFailedDelivery (orderItemIds string) orderentity.SetStatusToSOFFailedDeliveryResult {
    method := "/order/sof/failed_delivery"
    params := lib.InRow{
      "order_item_ids":orderItemIds,
    }
    result := orderentity.SetStatusToSOFFailedDeliveryResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
