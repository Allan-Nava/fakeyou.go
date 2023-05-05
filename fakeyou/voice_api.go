package fakeyou

import (
	"encoding/json"
	"fmt"
	"log"

	"gopkg.in/validator.v2"

	"github.com/Allan-Nava/fakeyou.go/constants/routes"
	"github.com/google/uuid"
)

/*
Get a list of voices
To download a list of all public voices, make a GET request to the following URL.

Note that the full model details are not available directly from the /tts/list API, and you'll have to consult the model details API to get usage statistics, etc
*/
func (f *fakeyou) GetListOfVoices() (*ResponseVoice, error) {
	//
	resp, err := f.restyGet(routes.LIST_VOICES, nil)
	if err != nil {
		return nil, err
	}
	var obj ResponseVoice
	if err := json.Unmarshal(resp.Body(), &obj); err != nil {
		return nil, err
	}
	return &obj, nil
}

/*
Get a list of voice categories
To download a list of voice categories (useful for building a voice search dropdown), use the following API:
*/

func (f *fakeyou) GetListOfVoiceCategories() (*ResponseVoiceCategories, error) {

	resp, err := f.restyGet(routes.CATEGORY_LIST_VOICES_TTS, nil)
	if err != nil {
		return nil, err
	}
	var obj ResponseVoiceCategories
	if err := json.Unmarshal(resp.Body(), &obj); err != nil {
		return nil, err
	}
	return &obj, nil
}

/*

Generate TTS Audio
Make a TTS request
To turn text into speech with your desired voice, you'll need to find the appropriate TTS model token from the model lookup API.

For example, TM:7wbtjphx8h8v in the following examples is our Mario * voice. (A paid voice actor that we hired to impersonate Mario).
*/

func (f *fakeyou) GenerateTTSAudio(text string, ttsModelToken string) (*ResponseGenerateTTS, error) {
	//
	body := &RequestGenerateTTS{
		TTSModelToken:        ttsModelToken,
		InferenceText:        text,
		UUIDIdempotencyToken: uuid.New().String(),
	}
	if errs := validator.Validate(body); errs != nil {
		// values not valid, deal with errors here
		return nil, errs
	}
	//
	resp, err := f.restyPost(routes.LOGIN, body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() == 429 {
		return nil, fmt.Errorf("IP IS BANNED (caused by login request)")
	}
	return nil, fmt.Errorf("")
}

/*
Poll TTS request status
Once you've submitted your TTS request, you'll want to poll for completion using the inference_job_token.
*/

func (f *fakeyou) PollTTSRequest(InferenceJobToken string) (*ResponsePollTTS, error) {
	//
	for {
		resp, err := f.restyGet(fmt.Sprintf(routes.POLL_TTS_REQUEST, InferenceJobToken), nil)
		if err != nil {
			return nil, err
		}
		log.Println("resp ", resp)
	} 
	return nil, nil
}
