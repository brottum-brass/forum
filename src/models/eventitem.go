package models

import "time"

type EventItem struct {
	ID                int       `json:"id"`
	Title             string    `json:"title"`
	Description       string    `json:"description"`
	InAssociationWith string    `json:"in_association_with,omitempty"`
	StartDate         time.Time `json:"start_date"`
	EndDate           time.Time `json:"end_date,omitempty"`
	Location          string    `json:"location"`
	IsFreeEntrance    bool      `json:"is_free_entrance"`
	TicketLink        string    `json:"ticket_link,omitempty"`
}

func (e *EventItem) IsUpcomingEvent() bool {
	return e.StartDate.After(time.Now())
}

func (e *EventItem) FormattedStartDate() string {
	return e.StartDate.Format("Jan 02")
}

func (e *EventItem) FormattedEndDate() string {
	return e.EndDate.Format("Jan 02")
}

func (e *EventItem) FormattedTime() string {
	return e.StartDate.Format("15:04")
}

func (e *EventItem) FormattedDateRange() string {
	if !e.EndDate.IsZero() && e.EndDate.Day() != e.StartDate.Day() {
		return e.FormattedStartDate() + " - " + e.FormattedEndDate()
	}

	return e.FormattedStartDate()
}
