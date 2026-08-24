package models

type BoardMemberItem struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email,omitempty"`
	Phone     string `json:"phone,omitempty"`
	ImagePath string `json:"image_path,omitempty"`
}
