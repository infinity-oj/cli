package clients

import (
	"fmt"
	"github.com/google/wire"
	"github.com/infinity-oj/server-v2/pkg/api"
	cookiejar "github.com/juju/persistent-cookiejar"
	"github.com/spf13/viper"
)

var Jar, _ = cookiejar.New(nil)

func NewClient(v *viper.Viper) api.API {
	client := api.New()

	client.SetCookieJar(Jar)
	client.SetHostUrl(fmt.Sprintf("%s/api/v1", v.Get("host").(string)))

	return client
}

var ProviderSet = wire.NewSet(
	NewClient,
)
