package fakeyou

type RequestLogin struct {
	Username string `json:"username_or_email" required:"true" validate:"nonnil,min=1"`
	Password string `json:"password" required:"true" validate:"nonnil,min=4"`
}
