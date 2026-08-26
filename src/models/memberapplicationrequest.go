package models

type MemberApplicationRequest struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	BirthYear      string `json:"birth_year"`
	Instrument     string `json:"instrument"`
	PreferredPart  string `json:"preferred_part"`
	Experience     string `json:"experience"`
	ReferenceName  string `json:"reference_name"`
	ReferencePhone string `json:"reference_phone"`
	Notes          string `json:"notes"`
}
