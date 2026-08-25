package models

type EventSubscriptionRequest struct {
	Email       string   `json:"email"`
	Preferences []string `json:"preferences"`
}
