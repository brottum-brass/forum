package models

type EventNotificationItem struct {
	Title                   string                   `json:"title"`
	Subtitle                string                   `json:"subtitle"`
	EmailLabel              string                   `json:"email_label"`
	EmailPlaceholder        string                   `json:"email_placeholder"`
	SubmitButtonLabel       string                   `json:"submit_button_label"`
	SuccessMessage          string                   `json:"success_message"`
	ErrorMessage            string                   `json:"error_message"`
	NotificationOptionItems []NotificationOptionItem `json:"notification_option_items"`
}
