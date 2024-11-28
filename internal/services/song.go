package services

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

// SongDetail представляет структуру данных для дополнительных данных о песне
type SongDetail struct {
	ReleaseDate string `json:"release_date"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

// GetSongDetails получает дополнительную информацию о песне
func GetSongDetails(group, title string) (*SongDetail, error) {
	url := fmt.Sprintf("https://api.example.com/song?group=%s&title=%s", group, title)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching song details: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	var songDetail SongDetail
	if err := json.NewDecoder(resp.Body).Decode(&songDetail); err != nil {
		log.Printf("Error decoding response: %v", err)
		return nil, err
	}

	return &songDetail, nil
}
