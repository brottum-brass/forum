package models

type ConductorItem struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Title     string `json:"title"`
	Bio       string `json:"bio"`
	ImagePath string `json:"image_path"`
	Vision    string `json:"vision"`
}
