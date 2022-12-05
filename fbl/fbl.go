package fbl
import(
	"github.com/wjpxxx/letgo/lib"
	lazadaConfig "github.com/wjpxxx/lazadago/config"
	fblentity "github.com/wjpxxx/lazadago/fbl/entity"
)
//Fbl
type Fbl struct{
	Config *lazadaConfig.Config
}

//GetPlatformProducts
//@Title Search products list
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=17&path=/fbl/platform_products/get
func (s *Fbl) GetPlatformProducts (perPage int,sellerId int64,marketplace string,sellerSku string,platformSkuName string,readyForInbound bool,platformSku string,page int) fblentity.GetPlatformProductsResult {
    method := "/fbl/platform_products/get"
    params := lib.InRow{
      "seller_id":sellerId,
      "marketplace":marketplace,
    }
    if perPage>-1{
        params["per_page"]=perPage
    }
    if sellerSku!=""{
        params["seller_sku"]=sellerSku
    }
    if platformSkuName!=""{
        params["platform_sku_name"]=platformSkuName
    }
    if platformSku!=""{
        params["platform_sku"]=platformSku
    }
    if page>-1{
        params["page"]=page
    }
    result := fblentity.GetPlatformProductsResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//GetFulfillmentProductDetail
//@Title GET  fulfillment product Detail；Call Get Platform Products for fulfillment_sku first
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=17&path=/fbl/fulfillment_products/get
func (s *Fbl) GetFulfillmentProductDetail (perPage int,shelfLifeFlag bool,marketplace string,fulfillmentSku string,serialNumberFlag bool,page int,fulfillmentSkuName string,barcode string) fblentity.GetFulfillmentProductDetailResult {
    method := "/fbl/fulfillment_products/get"
    params := lib.InRow{
      "marketplace":marketplace,
      "fulfillment_sku":fulfillmentSku,
    }
    if perPage>-1{
        params["per_page"]=perPage
    }
    if page>-1{
        params["page"]=page
    }
    if fulfillmentSkuName!=""{
        params["fulfillment_sku_name"]=fulfillmentSkuName
    }
    if barcode!=""{
        params["barcode"]=barcode
    }
    result := fblentity.GetFulfillmentProductDetailResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//GetInboundOrderDetail
//@Title Use this API to get the Inbound Order Detail
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=17&path=/fbl/inbound_order_detail/get
func (s *Fbl) GetInboundOrderDetail (inboundOrderNo string,marketplace string) fblentity.GetInboundOrderDetailResult {
    method := "/fbl/inbound_order_detail/get"
    params := lib.InRow{
      "inbound_order_no":inboundOrderNo,
      "marketplace":marketplace,
    }
    result := fblentity.GetInboundOrderDetailResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//GetInboundOrderList
//@Title Use this API to get inbound order list
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=17&path=/fbl/inbound_orders/get
func (s *Fbl) GetInboundOrderList (inboundOrderNo string,creationTimeFrom string,creationTimeTo string,inboundWarehouse string,sellerSku string,fulfillmentSku string,marketplace string,page string,perPage string) fblentity.GetInboundOrderListResult {
    method := "/fbl/inbound_orders/get"
    params := lib.InRow{
      "marketplace":marketplace,
    }
    if inboundOrderNo!=""{
        params["inbound_order_no"]=inboundOrderNo
    }
    if creationTimeFrom!=""{
        params["creation_time_From"]=creationTimeFrom
    }
    if creationTimeTo!=""{
        params["creation_time_To"]=creationTimeTo
    }
    if inboundWarehouse!=""{
        params["inbound_warehouse"]=inboundWarehouse
    }
    if sellerSku!=""{
        params["seller_sku"]=sellerSku
    }
    if fulfillmentSku!=""{
        params["fulfillment_sku"]=fulfillmentSku
    }
    if page!=""{
        params["page"]=page
    }
    if perPage!=""{
        params["per_page"]=perPage
    }
    result := fblentity.GetInboundOrderListResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//GetInventoryChangedSKU
//@Title Use this API to get SKU list
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=17&path=/fbl/inventory_changed_sku/get
func (s *Fbl) GetInventoryChangedSKU (warehouseCode string,page int,perPage int,marketPlace string,operateTimeFrom string,operateTimeTo string) fblentity.GetInventoryChangedSKUResult {
    method := "/fbl/inventory_changed_sku/get"
    params := lib.InRow{
      "market_place":marketPlace,
    }
    if warehouseCode!=""{
        params["warehouse_code"]=warehouseCode
    }
    if page>-1{
        params["page"]=page
    }
    if perPage>-1{
        params["per_page"]=perPage
    }
    if operateTimeFrom!=""{
        params["operate_Time_From"]=operateTimeFrom
    }
    if operateTimeTo!=""{
        params["operate_Time_To"]=operateTimeTo
    }
    result := fblentity.GetInventoryChangedSKUResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//GetInventoryOperateLog
//@Title Use this API to get a sku's inventory operate log
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=17&path=/fbl/inventory_operate_log/get
func (s *Fbl) GetInventoryOperateLog (page int,perPage int,marketPlace string,operateTimeFrom string,operateTimeTo string,warehouseCode string,fulfillmentSkuId string) fblentity.GetInventoryOperateLogResult {
    method := "/fbl/inventory_operate_log/get"
    params := lib.InRow{
      "market_place":marketPlace,
      "fulfillment_sku_id":fulfillmentSkuId,
    }
    if page>-1{
        params["page"]=page
    }
    if perPage>-1{
        params["per_page"]=perPage
    }
    if operateTimeFrom!=""{
        params["operate_time_from"]=operateTimeFrom
    }
    if operateTimeTo!=""{
        params["operate_time_to"]=operateTimeTo
    }
    if warehouseCode!=""{
        params["warehouse_code"]=warehouseCode
    }
    result := fblentity.GetInventoryOperateLogResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//GetOutboundOrderDetail
//@Title Use this API to Get outbound order detail; shoud call GetOutboundOrderList for outbound_order_no first
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=17&path=/fbl/outbound_order_detail/get
func (s *Fbl) GetOutboundOrderDetail (outboundOrderNo string,marketplace string) fblentity.GetOutboundOrderDetailResult {
    method := "/fbl/outbound_order_detail/get"
    params := lib.InRow{
      "outbound_order_no":outboundOrderNo,
      "marketplace":marketplace,
    }
    result := fblentity.GetOutboundOrderDetailResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//GetOutboundOrderList
//@Title Use this API to get outbound order list
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=17&path=/fbl/outbound_orders/get
func (s *Fbl) GetOutboundOrderList (outboundOrderNo string,creationTimeFrom string,creationTimeTo string,outboundWarehouse string,sellerSku string,fulfillmentSku string,marketplace string,page string,perPage string) fblentity.GetOutboundOrderListResult {
    method := "/fbl/outbound_orders/get"
    params := lib.InRow{
      "marketplace":marketplace,
    }
    if outboundOrderNo!=""{
        params["outbound_order_no"]=outboundOrderNo
    }
    if creationTimeFrom!=""{
        params["creation_time_from"]=creationTimeFrom
    }
    if creationTimeTo!=""{
        params["creation_time_to"]=creationTimeTo
    }
    if outboundWarehouse!=""{
        params["outbound_warehouse"]=outboundWarehouse
    }
    if sellerSku!=""{
        params["seller_sku"]=sellerSku
    }
    if fulfillmentSku!=""{
        params["fulfillment_sku"]=fulfillmentSku
    }
    if page!=""{
        params["page"]=page
    }
    if perPage!=""{
        params["per_page"]=perPage
    }
    result := fblentity.GetOutboundOrderListResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//GetWarehouseStock
//@Title Get SKU list and stock by warehouse code
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=17&path=/fbl/stocks/get
func (s *Fbl) GetWarehouseStock (sellerSku string,marketplace string,fulfilmentSku string,storeCode string) fblentity.GetWarehouseStockResult {
    method := "/fbl/stocks/get"
    params := lib.InRow{
      "marketplace":marketplace,
    }
    if sellerSku!=""{
        params["seller_sku"]=sellerSku
    }
    if fulfilmentSku!=""{
        params["fulfilment_sku"]=fulfilmentSku
    }
    if storeCode!=""{
        params["store_code"]=storeCode
    }
    result := fblentity.GetWarehouseStockResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//GetWarehouseStockV3
//@Title Get SKU list and stock by warehouse code, this version separates pending inbound and stock in transit in return json.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=17&path=/fbl/stocks/getV3
func (s *Fbl) GetWarehouseStockV3 (sellerSku string,marketplace string,fulfilmentSku string,storeCode string) fblentity.GetWarehouseStockV3Result {
    method := "/fbl/stocks/getV3"
    params := lib.InRow{
      "marketplace":marketplace,
    }
    if sellerSku!=""{
        params["seller_sku"]=sellerSku
    }
    if fulfilmentSku!=""{
        params["fulfilment_sku"]=fulfilmentSku
    }
    if storeCode!=""{
        params["store_code"]=storeCode
    }
    result := fblentity.GetWarehouseStockV3Result{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//UploadWaybill
//@Title Use this API to upload a waybill pdf to Lazada site. The maximum size of an pdf file is 1MB.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=17&path=/fbl/waybill/upload
func (s *Fbl) UploadWaybill (waybill []byte,packageCode string,trackingNumber string,extendsField string,storeCode string) fblentity.UploadWaybillResult {
    method := "/fbl/waybill/upload"
    params := lib.InRow{
      "@waybill":waybill,
      "package_code":packageCode,
      "tracking_number":trackingNumber,
      "store_code":storeCode,
    }
    if extendsField!=""{
        params["extends_field"]=extendsField
    }
    result := fblentity.UploadWaybillResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
