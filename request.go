package alexa

const (
	LocaleGerman      = "de-DE"
	LocaleEnAustralia = "en-AU"
	LocaleEnCanada    = "en-CA"
	LocaleEnUK        = "en-GB"
	LocaleEnIndia     = "en-IN"
	LocaleEnUS        = "en-US"
	LocaleFrench      = "fr-FR"
	LocaleJapanese    = "ja-JP"
)

type Request struct {
	Version     string      `json:"version"`
	Session     Session     `json:"session"`
	Context     Context     `json:"context"`
	RequestBody RequestBody `json:"request"`
}

type Session struct {
	New         bool                   `json:"new"`
	SessionId   string                 `json:"sessionId"`
	Attributes  map[string]interface{} `json:"attributes"`
	Application struct {
		ApplicationId string `json:"applicationId"`
	} `json:"application"`
	User struct {
		UserId      string `json:"userId"`
		AccessToken string `json:"accessToken,omitempty"`
		Permissions struct {
			ConsentToken string `json:"consentToken,omitempty"`
		} `json:"permissions,omitempty"`
	} `json:"user"`
}

type Context struct {
	System      System      `json:"system"`
	AudioPlayer AudioPlayer `json:"audioPlayer"`
}

type System struct {
	ApiAccessToken string `json:"apiAccessToken"`
	ApiEndpoint    string `json:"apiEndpoint"`
	Application    struct {
		ApplicationId string `json:"applicationId"`
	} `json:"application"`
	Device struct {
		DeviceId            string              `json:"deviceId"`
		SupportedInterfaces SupportedInterfaces `json:"supportedInterfaces,omitempty"`
	} `json:"device"`
	User struct {
		UserId      string `json:"userId"`
		AccessToken string `json:"accessToken,omitempty"`
		Permissions struct {
			ConsentToken string `json:"consentToken,omitempty"`
		} `json:"permissions,omitempty"`
	} `json:"user"`
}

type SupportedInterfaces struct {
	Display     Display     `json:"display,omitempty"`
	AudioPlayer AudioPlayer `json:"audioPlayer,omitempty"`
	VideoApp    VideoApp    `json:"videoApp,omitempty"`
}

type Display struct {
	TemplateVersion string `json:"templateVersion,omitempty"`
	MarkupVersion   string `json:"markupVersion,omitempty"`
}

type VideoApp struct {
}

type AudioPlayer struct {
	Token                string `json:"token"`
	OffsetInMilliseconds int64  `json:"offsetInMilliseconds"`
	PlayerActivity       string `json:"playerActivity"`
}

type RequestBody struct {
	Type        string `json:"type"`
	Timestamp   string `json:"timestamp"`
	RequestId   string `json:"requestId"`
	Locale      string `json:"locale"`
	DialogState string `json:"dialogState,omitempty"`
	Intent      Intent `json:"intent,omitempty"`
	Reason      string `json:"reason,omitempty"`
	Error       struct {
		Type    string `json:"type"`
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

type Intent struct {
	Name               string          `json:"name"`
	ConfirmationStatus string          `json:"confirmationStatus"`
	Slots              map[string]Slot `json:"slots"`
}

type Slot struct {
	Name               string      `json:"name"`
	Value              string      `json:"value"`
	ConfirmationStatus string      `json:"confirmationStatus"`
	Resolutions        interface{} `json:"resolutions,omitempty"`
}
