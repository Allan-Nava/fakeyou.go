package configuration

import (
	"github.com/caarlos0/env/v6"
	"github.com/go-resty/resty/v2"

	"github.com/Allan-Nava/fakeyou.go/constants/routes"
)

type Configuration struct {
	IsDebug    bool `env:"IS_DEBUG"`
	BaseUrl    string
	RestClient *resty.Client
}

func GetConfiguration() *Configuration {
	configuration := Configuration{}
	err := env.Parse(&configuration)
	if err != nil {
		panic("failed to read configuration")
	}
	//
	configuration.BaseUrl = routes.BASE_URL
	configuration.RestClient = resty.New()
	configuration.RestClient.SetHeader("Content-Type", "application/json")
	if configuration.IsDebug {
		configuration.RestClient.SetDebug(true)
	}
	//
	return &configuration
}
