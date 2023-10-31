package fakeyou

import (
	"github.com/go-resty/resty/v2"

	"github.com/Allan-Nava/fakeyou.go/configuration"
	"github.com/Allan-Nava/fakeyou.go/constants/routes"
)

type IFakeYou interface {
	//
	IsDebug() bool
	// voice api
	GetListOfVoices() (*ResponseVoice, error)
	GetListOfVoiceCategories() (*ResponseVoiceCategories, error)
	GenerateTTSAudio(text string, ttsModelToken string) (*ResponseGenerateTTS, error)
	PollTTSRequest(InferenceJobToken string) (*ResponsePollTTS, error)
	// auth api
	Login(body RequestLogin) error
	//
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
	fk.restClient.SetBaseURL(routes.BASE_URL)
	fk.restClient.SetHeader("Content-Type", "application/json")
	if configuration.isDebug {
		fk.restClient.SetDebug(true)
	}
	return fk
}
//