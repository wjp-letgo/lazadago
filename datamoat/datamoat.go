package datamoat
import(
	"github.com/wjpxxx/letgo/lib"
	lazadaConfig "github.com/wjpxxx/lazadago/config"
	datamoatentity "github.com/wjpxxx/lazadago/datamoat/entity"
)
//DataMoat
type DataMoat struct{
	Config *lazadaConfig.Config
}

//DataMoatComputeRisk
//@Title Note that currently all the regions must use the domain "api.lazada.com" to invoke this API. 
//This API is used to access DataMoat Account Security Service, which is required in the process of accessing sensitive data.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=15&path=/datamoat/compute_risk
func (s *DataMoat) DataMoatComputeRisk (time string,appName string,userId string,userIp string,ati string) datamoatentity.DataMoatComputeRiskResult {
    method := "/datamoat/compute_risk"
    params := lib.InRow{
      "time":time,
      "appName":appName,
      "userId":userId,
      "userIp":userIp,
      "ati":ati,
    }
    result := datamoatentity.DataMoatComputeRiskResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
//DataMoatLogin
//@Title Note that currently all the regions must use the domain "api.lazada.com" to invoke this API. 
//This API is used to access DataMoat Account Security Service, which is required in the process of accessing sensitive data.
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.78de6bbeqtnjmQ#/api?cid=15&path=/datamoat/login
func (s *DataMoat) DataMoatLogin (time string,appName string,userId string,tid string,userIp string,ati string,loginResult string,loginMessage string) datamoatentity.DataMoatLoginResult {
    method := "/datamoat/login"
    params := lib.InRow{
      "time":time,
      "appName":appName,
      "userId":userId,
      "tid":tid,
      "userIp":userIp,
      "ati":ati,
      "loginResult":loginResult,
      "loginMessage":loginMessage,
    }
    result := datamoatentity.DataMoatLoginResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
