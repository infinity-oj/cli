package clients

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/google/wire"
	"github.com/infinity-oj/cli/internal/clients/accounts"
	"github.com/spf13/viper"
)

// Options is log configuration struct
type Options struct {
	Url string `yaml:"url"`
}

func NewOptions(v *viper.Viper) (*Options, error) {
	var (
		err error
		o   = new(Options)
	)
	if err = v.UnmarshalKey("http", o); err != nil {
		return nil, err
	}

	fmt.Println(o)

	return o, err
}

func NewClient(options *Options) *resty.Client {
	client := resty.New()
	client.SetHostURL(options.Url)
	return client
}

var ProviderSet = wire.NewSet(NewClient, NewOptions, accounts.NewAccountClient)
