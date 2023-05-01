package routes

///

const (
	//
	BASE_URL = "https://api.fakeyou.com/"
	// routes api
	LIST_VOICES              = "tts/list"
	CATEGORY_LIST_VOICES_TTS = "category/list/tts"
	GENERATE_TTS_AUDIO       = "tts/inference"
	POLL_TTS_REQUEST         = "tts/job/%s"
)
