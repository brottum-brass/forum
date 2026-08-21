package theme

import "context"

type Mode struct {
	Name                 string `json:"name"`
	PrimaryBackground    string `json:"primary_background"`
	SecondaryBackground  string `json:"secondary_background"`
	PrimaryText          string `json:"primary_text"`
	SecondaryText        string `json:"secondary_text"`
	PrimaryBorder        string `json:"primary_border"`
	SecondaryBorder      string `json:"secondary_border"`
	AccentText           string `json:"accent_text"`
	AccentHoverText      string `json:"accent_hover_text"`
	AccentGroupHoverText string `json:"accent_group_hover_text"`
	AccentFocusRing      string `json:"accent_focus_ring"`
	AccentGradient       string `json:"accent_gradient"`
	HoverBackground      string `json:"hover_background"`
	HoverBorder          string `json:"hover_border"`
	IconFilter           string `json:"icon_filter"`
}

type Theme interface {
	LoadThemes() error
	T(ctx context.Context, key string) Mode
}
