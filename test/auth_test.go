package test

import (
	"testing"

	"github.com/Allan-Nava/fakeyou.go/fakeyou"
)

func Test_Login(t *testing.T) {
	f := GetFakeYou()
	//expected :=

	body := fakeyou.RequestLogin{
		Username: "wokeye4880@saeoil.com",
		Password: "fakeyou.go",
	}
	err := f.Login(body)
	if err != nil {
		panic(err)
	}
	//
}
