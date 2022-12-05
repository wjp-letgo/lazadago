package seller
import(
	"github.com/wjpxxx/letgo/lib"
	lazadaConfig "github.com/wjpxxx/lazadago/config"
	sellerentity "github.com/wjpxxx/lazadago/seller/entity"
)
//Seller
type Seller struct{
	Config *lazadaConfig.Config
}

//GetMultiWarehouseBySeller
//@Title Provide seller multi warehouse address data of the specific seller, like warehouse code and warehouse name and etc.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=2&path=/seller/warehouse/get
func (s *Seller) GetMultiWarehouseBySeller (addressTypes []string) sellerentity.GetMultiWarehouseBySellerResult {
    method := "/seller/warehouse/get"
    params := lib.InRow{
      "addressTypes":addressTypes,
    }
    result := sellerentity.GetMultiWarehouseBySellerResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//GetSeller
//@Title Get seller information by current seller ID.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=2&path=/seller/get
func (s *Seller) GetSeller () sellerentity.GetSellerResult {
    method := "/seller/get"
    params := lib.InRow{}
    result := sellerentity.GetSellerResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//GetSellerMetricsById
//@Title Provide seller metrics data of the specific seller, like positive seller rating, ship on time rate and etc.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=2&path=/seller/metrics/get
func (s *Seller) GetSellerMetricsById () sellerentity.GetSellerMetricsByIdResult {
    method := "/seller/metrics/get"
    params := lib.InRow{}
    result := sellerentity.GetSellerMetricsByIdResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//GetSellerPerformance
//@Title Provide the performance metrics of the current seller, such as positive seller rating, ship on time, etc.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=2&path=/seller/performance/get
func (s *Seller) GetSellerPerformance (language string) sellerentity.GetSellerPerformanceResult {
    method := "/seller/performance/get"
    params := lib.InRow{
    }
    if language!=""{
        params["language"]=language
    }
    result := sellerentity.GetSellerPerformanceResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//UpdateSeller
//@Title Use this API to update the email address of the seller who makes the call.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=2&path=/seller/update
func (s *Seller) UpdateSeller (payload string) sellerentity.UpdateSellerResult {
    method := "/seller/update"
    params := lib.InRow{
      "payload":payload,
    }
    result := sellerentity.UpdateSellerResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//UpdateUser
//@Title Use this API to update the email address of a user under the seller account.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=2&path=/user/update
func (s *Seller) UpdateUser (payload string) sellerentity.UpdateUserResult {
    method := "/user/update"
    params := lib.InRow{
      "payload":payload,
    }
    result := sellerentity.UpdateUserResult{}
    err := s.Config.HttpPost(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
