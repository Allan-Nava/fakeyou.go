package fakeyou

type RequestLogin struct {
	Username string `json:"username_or_email" required:"true" validate:"nonnil,min=1"`
	Password string `json:"password" required:"true" validate:"nonnil,min=4"`
}

type RequestGenerateTTS struct {
	TTSModelToken        string `json:"tts_model_token" required:"true" validate:"nonnil,min=4"`
	UUIDIdempotencyToken string `json:"uuid_idempotency_token" required:"true" validate:"nonnil,min=4"`
	InferenceText        string `json:"inference_text" required:"true" validate:"nonnil,min=4"`
}
