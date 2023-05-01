package fakeyou

import (
	"github.com/go-resty/resty/v2"

	"github.com/Allan-Nava/fakeyou.go/configuration"
	"github.com/Allan-Nava/fakeyou.go/constants/routes"
)

type IFakeYou interface {
	// voice api
	GetListOfVoices()
	// auth api
	Login(body RequestLogin) error
}

type fakeyou struct {
	configuration *configuration.Configuration
	restClient    *resty.Client
}

func NewFakeYou(configuration *configuration.Configuration) IFakeYou {
	fk := &fakeyou{
		configuration: configuration,
	}
	fk.restClient = resty.New()
	fk.restClient.SetHostURL(routes.BASE_URL)
	fk.restClient.SetHeader("Content-Type", "application/json")
	if configuration.IsDebug {
		fk.restClient.SetDebug(true)
	}
	return fk
}
