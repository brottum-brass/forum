package models

type EventSubscriptionRequestItem struct {
	Email       string   `json:"email"`
	Preferences []string `json:"preferences"`
}
