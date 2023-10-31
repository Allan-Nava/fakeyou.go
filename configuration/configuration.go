package configuration

import (
	"github.com/caarlos0/env/v6"

	"github.com/Allan-Nava/fakeyou.go/constants/routes"
)

type Configuration struct {
	Debug bool `env:"IS_DEBUG"`
	BaseUrl string
	//RestClient *resty.Client
}

func GetConfiguration() *Configuration {
	configuration := Configuration{}
	err := env.Parse(&configuration)
	if err != nil {
		panic("failed to read configuration")
	}
	//
	configuration.BaseUrl = routes.BASE_URL
	//
	return &configuration
}
