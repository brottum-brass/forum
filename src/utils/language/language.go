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
	Events   Events  `json:"events"`
	About    About   `json:"about"`
	Members  Members `json:"members"`
	NotFound struct {
		Title    string `json:"title"`
		Message  string `json:"message"`
		HomeLink string `json:"home_link"`
	} `json:"not_found"`
}

type About struct {
	PageTitle string `json:"page_title"`
	Subtitle  string `json:"subtitle"`
	Stats     struct {
		EstablishedLabel string `json:"established_label"`
		EstablishedYear  string `json:"established_year"`
		MembersLabel     string `json:"members_label"`
		MembersCount     string `json:"members_count"`
		DivisionLabel    string `json:"division_label"`
		DivisionName     string `json:"division_name"`
		LocationLabel    string `json:"location_label"`
		Location         string `json:"location"`
	} `json:"stats"`
	History struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	} `json:"history"`
	Mission struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	} `json:"mission"`
	Conductor struct {
		TitleSingular  string                 `json:"title_singular"`
		TitlePlural    string                 `json:"title_plural"`
		VisionLabel    string                 `json:"vision_label"`
		ConductorItems []models.ConductorItem `json:"conductor_items"`
	} `json:"conductor"`
	Competition struct {
		Title                  string                         `json:"title"`
		YearLabel              string                         `json:"year_label"`
		EventLabel             string                         `json:"event_label"`
		PieceLabel             string                         `json:"piece_label"`
		ConductorLabel         string                         `json:"conductor_label"`
		PositionLabel          string                         `json:"position_label"`
		CompetitionResultItems []models.CompetitionResultItem `json:"competition_result_items"`
	} `json:"competition"`
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
