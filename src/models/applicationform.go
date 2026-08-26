package models

type ApplicationForm struct {
	Title                     string `json:"title"`
	Subtitle                  string `json:"subtitle"`
	NameLabel                 string `json:"name_label"`
	NamePlaceholder           string `json:"name_placeholder"`
	EmailLabel                string `json:"email_label"`
	EmailPlaceholder          string `json:"email_placeholder"`
	PhoneLabel                string `json:"phone_label"`
	PhonePlaceholder          string `json:"phone_placeholder"`
	BirthYearLabel            string `json:"birth_year_label"`
	BirthYearPlaceholder      string `json:"birth_year_placeholder"`
	InstrumentLabel           string `json:"instrument_label"`
	InstrumentPlaceholder     string `json:"instrument_placeholder"`
	PreferredPartLabel        string `json:"preferred_part_label"`
	PreferredPartPlaceholder  string `json:"preferred_part_placeholder"`
	ExperienceLabel           string `json:"experience_label"`
	ExperiencePlaceholder     string `json:"experience_placeholder"`
	ReferenceNameLabel        string `json:"reference_name_label"`
	ReferenceNamePlaceholder  string `json:"reference_name_placeholder"`
	ReferencePhoneLabel       string `json:"reference_phone_label"`
	ReferencePhonePlaceholder string `json:"reference_phone_placeholder"`
	NotesLabel                string `json:"notes_label"`
	NotesPlaceholder          string `json:"notes_placeholder"`
	SubmitButtonLabel         string `json:"submit_button_label"`
	SuccessMessage            string `json:"success_message"`
	ErrorMessage              string `json:"error_message"`
}
