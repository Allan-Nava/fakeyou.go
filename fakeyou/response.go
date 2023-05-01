package fakeyou

type BaseResponse struct {
	Success bool `json:"success" `
}

type Model struct {
	ModelToken                     string               `json:"model_token"`
	TTSModelType                   TTSModelType         `json:"tts_model_type"`
	CreatorUserToken               string               `json:"creator_user_token"`
	CreatorUsername                string               `json:"creator_username"`
	CreatorDisplayName             string               `json:"creator_display_name"`
	CreatorGravatarHash            string               `json:"creator_gravatar_hash"`
	Title                          string               `json:"title"`
	IETFLanguageTag                IETFTag              `json:"ietf_language_tag"`
	IETFPrimaryLanguageSubtag      IETFTag              `json:"ietf_primary_language_subtag"`
	IsFrontPageFeatured            bool                 `json:"is_front_page_featured"`
	IsTwitchFeatured               bool                 `json:"is_twitch_featured"`
	MaybeSuggestedUniqueBotCommand *string              `json:"maybe_suggested_unique_bot_command"`
	CreatorSetVisibility           CreatorSetVisibility `json:"creator_set_visibility"`
	UserRatings                    UserRatings          `json:"user_ratings"`
	CategoryTokens                 []string             `json:"category_tokens"`
	CreatedAt                      string               `json:"created_at"`
	UpdatedAt                      string               `json:"updated_at"`
}

type UserRatings struct {
	PositiveCount int64 `json:"positive_count"`
	NegativeCount int64 `json:"negative_count"`
	TotalCount    int64 `json:"total_count"`
}

type CreatorSetVisibility string

const (
	Public CreatorSetVisibility = "public"
)

type IETFTag string

const (
	Ar    IETFTag = "ar"
	En    IETFTag = "en"
	EnUS  IETFTag = "en-US"
	Es    IETFTag = "es"
	Es419 IETFTag = "es-419"
)

type TTSModelType string

const (
	Tacotron2 TTSModelType = "tacotron2"
)

type ResponseVoice struct {
	BaseResponse
	Models []Model `json:"models"`
}
