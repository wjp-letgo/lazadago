package logistics

import(
	"github.com/wjpxxx/letgo/lib"
	lazadaConfig "github.com/wjpxxx/lazadago/config"
	logisticsentity "github.com/wjpxxx/lazadago/logistics/entity"
)
//Logistics
type Logistics struct{
	Config *lazadaConfig.Config
}

//GetOrderTrace
//@Title Query logistic detail for seller erp with seller id, order id and locale info. This api is only available in the state after ready to ship.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=6&path=/logistic/order/trace
func (s *Logistics) GetOrderTrace (sellerId string,orderId string,locale string,ofcPackageIdList []string) logisticsentity.GetOrderTraceResult {
    method := "/logistic/order/trace"
    params := lib.InRow{
      "seller_id":sellerId,
      "order_id":orderId,
    }
    if locale!=""{
        params["locale"]=locale
    }
    result := logisticsentity.GetOrderTraceResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//GetShipmentProviders
//@Title Use this API to get the list of all active shipping providers, which is needed when working with the SetStatusToPackedByMarketplace API.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=6&path=/shipment/providers/get
func (s *Logistics) GetShipmentProviders () logisticsentity.GetShipmentProvidersResult {
    method := "/shipment/providers/get"
    params := lib.InRow{}
    result := logisticsentity.GetShipmentProvidersResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
