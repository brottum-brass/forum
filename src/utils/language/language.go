package language

import "context"

type Content struct {
	Common struct {
		Language string `json:"language"`
		Title    string `json:"title"`
		Name     string `json:"name"`
	} `json:"common"`
	Header struct {
		Home    string `json:"home"`
		Events  string `json:"events"`
		About   string `json:"about"`
		Members string `json:"members"`
		Contact string `json:"contact"`
	} `json:"header"`
	Footer struct {
		Copyright string `json:"copyright"`
	} `json:"footer"`
}

type Language interface {
	LoadLanguages() error
	L(ctx context.Context, key string) Content
}
