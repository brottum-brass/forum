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
	Contact  Contact `json:"contact"`
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
		TitleSingular string             `json:"title_singular"`
		TitlePlural   string             `json:"title_plural"`
		VisionLabel   string             `json:"vision_label"`
		Conductors    []models.Conductor `json:"conductors"`
	} `json:"conductor"`
	Competition struct {
		Title              string                     `json:"title"`
		YearLabel          string                     `json:"year_label"`
		EventLabel         string                     `json:"event_label"`
		PieceLabel         string                     `json:"piece_label"`
		ConductorLabel     string                     `json:"conductor_label"`
		PositionLabel      string                     `json:"position_label"`
		CompetitionResults []models.CompetitionResult `json:"competition_results"`
	} `json:"competition"`
}

type Events struct {
	PageTitle          string                   `json:"page_title"`
	SwipeHint          string                   `json:"swipe_hint"`
	NoEvents           string                   `json:"no_events"`
	GetTickets         string                   `json:"get_tickets"`
	TicketsUnavailable string                   `json:"tickets_unavailable"`
	FreeEntrance       string                   `json:"free_entrance"`
	EventItems         []models.Event           `json:"event_items"`
	Notification       models.EventNotification `json:"notification"`
}

func (e *Events) GetEventByID(id int) (*models.Event, error) {
	for _, event := range e.EventItems {
		if event.ID == id {
			return &event, nil
		}
	}

	return nil, fmt.Errorf("event with ID %d not found", id)
}

type Members struct {
	PageTitle   string          `json:"page_title"`
	Subtitle    string          `json:"subtitle"`
	NoMembers   string          `json:"no_members"`
	MemberItems []models.Member `json:"member_items"`
}

func (m *Members) GetMemberByID(id int) (*models.Member, error) {
	for _, member := range m.MemberItems {
		if member.ID == id {
			return &member, nil
		}
	}

	return nil, fmt.Errorf("member with ID %d not found", id)
}

type Contact struct {
	PageTitle    string             `json:"page_title"`
	Subtitle     string             `json:"subtitle"`
	BoardSection BoardSection       `json:"board_section"`
	InquiryForm  models.InquiryForm `json:"inquiry_form"`
}

type BoardSection struct {
	Title          string               `json:"title"`
	Subtitle       string               `json:"subtitle"`
	NoBoardMembers string               `json:"no_board_members"`
	BoardMembers   []models.BoardMember `json:"board_members"`
}
