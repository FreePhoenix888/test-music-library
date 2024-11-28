package models

// Song представляет структуру песни
type Song struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Group       string `json:"group"`
	Title       string `json:"title"`
	ReleaseDate string `json:"release_date"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

