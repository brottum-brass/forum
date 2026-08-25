package models

type InquiryForm struct {
	Title              string   `json:"title"`
	Subtitle           string   `json:"subtitle"`
	NameLabel          string   `json:"name_label"`
	NamePlaceholder    string   `json:"name_placeholder"`
	EmailLabel         string   `json:"email_label"`
	EmailPlaceholder   string   `json:"email_placeholder"`
	PhoneLabel         string   `json:"phone_label"`
	PhonePlaceholder   string   `json:"phone_placeholder"`
	TypeLabel          string   `json:"type_label"`
	MessageLabel       string   `json:"message_label"`
	MessagePlaceholder string   `json:"message_placeholder"`
	SubmitButtonLabel  string   `json:"submit_button_label"`
	SuccessMessage     string   `json:"success_message"`
	ErrorMessage       string   `json:"error_message"`
	Options            []Option `json:"options"`
}
