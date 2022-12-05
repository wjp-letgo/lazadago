package system

import(
	"github.com/wjpxxx/letgo/lib"
	lazadaConfig "github.com/wjpxxx/lazadago/config"
	systementity "github.com/wjpxxx/lazadago/system/entity"
)
//System
type System struct{
	Config *lazadaConfig.Config
}
//GenerateAccessToken
//@Title generate access_token for call api
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.bd786bbeNJUDaJ#/api?cid=11&path=/auth/token/create
func (s *System) GenerateAccessToken (code string,uuid string) systementity.GenerateAccessTokenResult {
    method := "/auth/token/create"
    params := lib.InRow{
      "code":code,
    }
    if uuid!=""{
        params["uuid"]=uuid
    }
    result := systementity.GenerateAccessTokenResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}

//RefreshAccessToken
//@Title refresh access_token, the endpoint is https://auth.lazada.com/rest
//@Description https://open.lazada.com/doc/api.htm?spm=a2o9m.11193531.0.0.bd786bbeNJUDaJ#/api?cid=11&path=/auth/token/refresh
func (s *System) RefreshAccessToken (refreshToken string) systementity.RefreshAccessTokenResult {
    method := "/auth/token/refresh"
    params := lib.InRow{
      "refresh_token":refreshToken,
    }
    result := systementity.RefreshAccessTokenResult{}
    err := s.Config.HttpGet(method, params, &result)
    if err != nil {
        result.Code = err.Error()
    }
    return result
}
