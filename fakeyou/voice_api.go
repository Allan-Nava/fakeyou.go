package fakeyou

import (
	"encoding/json"

	"github.com/Allan-Nava/fakeyou.go/constants/routes"
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
