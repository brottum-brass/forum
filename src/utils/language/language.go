package language

import (
	"context"

	"github.com/brottum-brass/forum/src/models"
)

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
	Events struct {
		PageTitle  string             `json:"page_title"`
		SwipeHint  string             `json:"swipe_hint"`
		NoEvents   string             `json:"no_events"`
		GetTickets string             `json:"get_tickets"`
		EventItems []models.EventItem `json:"event_items"`
	} `json:"events"`
}

type Language interface {
	LoadLanguages() error
	L(ctx context.Context, key string) Content
}
