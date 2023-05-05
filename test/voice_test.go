package test

import (
	"log"
	"testing"
)

func Test_ListVoices(t *testing.T) {
	f := GetFakeYou()
	resp, err := f.GetListOfVoices()
	if err != nil {
		panic(err)
	}
	log.Println("resp", resp)
}

func Test_ListVoiceCategories(t *testing.T) {
	f := GetFakeYou()
	resp, err := f.GetListOfVoiceCategories()
	if err != nil {
		panic(err)
	}
	log.Println("resp", resp)
}

func Test_GenerateTTSAudio(t *testing.T) {
	f := GetFakeYou()
	resp, err := f.GenerateTTSAudio("Hello World", "TM:7wbtjphx8h8v")
	if err != nil {
		panic(err)
	}
	log.Println("resp", resp)
}

func Test_PollTTSRequest(t *testing.T) {
	f := GetFakeYou()
	resp, err := f.GenerateTTSAudio("Hello World", "TM:7wbtjphx8h8v")
	if err != nil {
		panic(err)
	}
	log.Println("resp", resp)
	pollResp, err := f.PollTTSRequest(resp.InferenceJobToken)
	if err != nil {
		panic(err)
	}
	log.Println("pollResp", pollResp)
}
