package fakeyou

import (

	"github.com/go-resty/resty/v2"
)

type IBuilder interface {
	Build() []string
}

type Builder fakeyou

func (f *Builder) Build() []string {
	var args []string
	return args
}


func (f *fakeyou) IsDebug() bool {
	return f.configuration.IsDebug
}

// Resty Methods

func (f *fakeyou) restyPost(url string, body interface{}) (*resty.Response, error) {
	resp, err := f.restClient.R().
		SetHeader("Accept", "application/json").
		SetBody(body).
		Post(url)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *fakeyou) restyGet(url string, queryParams map[string]string) (*resty.Response, error) {
	resp, err := o.restClient.R().
		SetQueryParams(queryParams).
		Get(url)
	//
	if err != nil {
		return nil, err
	}
	return resp, nil
}
