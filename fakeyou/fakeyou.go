package fakeyou

import (
	"github.com/Allan-Nava/fakeyou.go/configuration"
)

type IFakeYou interface {
}

type fakeyou struct {
	configuration *configuration.Configuration
}

func NewFakeYou(configuration *configuration.Configuration) IFakeYou {
	return &fakeyou{
		configuration: configuration,
	}
}
