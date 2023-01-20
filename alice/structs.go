package alice

type Request struct {
	Meta struct {
		Locale     string `json:"locale"`
		Timezone   string `json:"timezone"`
		ClientId   string `json:"client_id"`
		Interfaces struct {
			Screen struct {
			} `json:"screen"`
			Payments struct {
			} `json:"payments"`
			AccountLinking struct {
			} `json:"account_linking"`
		} `json:"interfaces"`
	} `json:"meta"`
	Session struct {
		MessageId int    `json:"message_id"`
		SessionId string `json:"session_id"`
		SkillId   string `json:"skill_id"`
		User      struct {
			UserId string `json:"user_id"`
		} `json:"user"`
		Application struct {
			ApplicationId string `json:"application_id"`
		} `json:"application"`
		New    bool   `json:"new"`
		UserId string `json:"user_id"`
	} `json:"session"`
	Request struct {
		Command           string `json:"command"`
		OriginalUtterance string `json:"original_utterance"`
		Nlu               struct {
			Tokens   []string      `json:"tokens"`
			Entities []interface{} `json:"entities"`
			Intents  struct {
			} `json:"intents"`
		} `json:"nlu"`
		Markup struct {
			DangerousContext bool `json:"dangerous_context"`
		} `json:"markup"`
		Type string `json:"type"`
	} `json:"request"`
	Version string `json:"version"`
}

type Resp struct {
	Text       string `json:"text"`
	EndSession bool   `json:"end_session"`
}

type Response struct {
	Response Resp   `json:"response"`
	Version  string `json:"version"`
}
