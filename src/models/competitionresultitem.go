package models

type CompetitionResultItem struct {
	ID        int    `json:"id"`
	Year      int    `json:"year"`
	Event     string `json:"event"`
	Piece     string `json:"piece"`
	Conductor string `json:"conductor"`
	Position  string `json:"position"`
}
