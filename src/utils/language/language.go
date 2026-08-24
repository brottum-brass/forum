package language

import (
	"context"
	"fmt"

	"github.com/brottum-brass/forum/src/models"
)

type Language interface {
	LoadLanguages() error
	L(ctx context.Context, key string) Content
}

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
	Events  Events  `json:"events"`
	Members Members `json:"members"`
}

type Events struct {
	PageTitle  string             `json:"page_title"`
	SwipeHint  string             `json:"swipe_hint"`
	NoEvents   string             `json:"no_events"`
	GetTickets string             `json:"get_tickets"`
	EventItems []models.EventItem `json:"event_items"`
}

func (e *Events) GetEventByID(id int) (*models.EventItem, error) {
	for _, event := range e.EventItems {
		if event.ID == id {
			return &event, nil
		}
	}

	return nil, fmt.Errorf("not found")
}

type Members struct {
	PageTitle   string              `json:"page_title"`
	Subtitle    string              `json:"subtitle"`
	NoMembers   string              `json:"no_members"`
	MemberItems []models.MemberItem `json:"member_items"`
}

func (m *Members) GetMemberByID(id int) (*models.MemberItem, error) {
	for _, member := range m.MemberItems {
		if member.ID == id {
			return &member, nil
		}
	}

	return nil, fmt.Errorf("not found")
}
