package fakeyou

import (
	"fmt"
	"log"

	"gopkg.in/validator.v2"

	"github.com/Allan-Nava/fakeyou.go/constants/routes"
)

// Authentication API

func (f *fakeyou) Login(body RequestLogin) error {
	//
	if errs := validator.Validate(body); errs != nil {
		// values not valid, deal with errors here
		return errs
	}
	//
	resp, err := f.restyPost(routes.LOGIN, body)
	if err != nil {
		return err
	}
	if resp.StatusCode() == 429 {
		return fmt.Errorf("IP IS BANNED (caused by login request)")
	}
	if resp.IsError() {
		//login error
		//return fmt.Errorf( resp.RawResponse.)
	}

	log.Println("Processing the response (login)")

	return nil
}
