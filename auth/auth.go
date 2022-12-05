package auth
import (
	"fmt"

	lazadaConfig "github.com/wjpxxx/lazadago/config"
)

//Auth
type Auth struct {
	Config *lazadaConfig.Config
}

//AuthorizationURL
func (a *Auth) AuthorizationURL(redirectUri string) string {
	return fmt.Sprintf("%s?response_type=code&force_auth=true&redirect_uri=%s&client_id=%s",a.Config.GetApiUrl("all"),redirectUri,a.Config.AppKey)
}