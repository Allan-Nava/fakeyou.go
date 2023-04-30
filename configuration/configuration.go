package configuration

import (
	"github.com/caarlos0/env/v6"
	"github.com/go-resty/resty/v2"

	"github.com/Allan-Nava/fakeyou.go/constants/routes"
)

type Configuration struct {
	IsDebug    bool `env:"IS_DEBUG"`
	BaseUrl    string
	restClient *resty.Client
}

func GetConfiguration() *Configuration {
	configuration := Configuration{}
	err := env.Parse(&configuration)
	if err != nil {
		panic("failed to read configuration")
	}
	//
	configuration.BaseUrl = routes.BASE_URL
	configuration.restClient = resty.New()
	configuration.restClient.SetHeader("Content-Type", "application/json")
	if configuration.IsDebug {
		configuration.restClient.SetDebug(true)
	}
	//
	return &configuration
}
