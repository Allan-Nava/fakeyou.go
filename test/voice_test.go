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
