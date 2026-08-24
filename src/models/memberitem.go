package models

type MemberItem struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Bio           string `json:"bio"`
	ImagePath     string `json:"image_path"`
	Section       string `json:"section"`
	Instrument    string `json:"instrument"`
	FavoritePiece string `json:"favorite_piece"`
}
