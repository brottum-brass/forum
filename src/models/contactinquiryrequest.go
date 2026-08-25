package models

type ContactInquiryRequest struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	InquiryType string `json:"inquiry_type"`
	Message     string `json:"message"`
}
