package test

import (
	"os"
	"testing"

	"github.com/Allan-Nava/fakeyou.go/configuration"
	"github.com/Allan-Nava/fakeyou.go/env"
	"github.com/Allan-Nava/fakeyou.go/fakeyou"
)

func TestMain(m *testing.M) {
	if os.Getenv("APP_ENV") == "" {
		err := os.Setenv("APP_ENV", "test")
		if err != nil {
			panic("could not set test env")
		}
	}
	env.Load()
	m.Run()
}

func GetFakeYou() fakeyou.IFakeYou {
	//
	configuration := configuration.GetConfiguration()
	f := fakeyou.NewFakeYou(
		configuration,
	)
	return f
}

//
