package handlers

import (
	"encoding/json"
	"go-music-library/internal/models"
	"go-music-library/internal/services"
	"log"
	"net/http"
	"strconv"
	"gorm.io/gorm"

)

// GetSongs обрабатывает запрос на получение списка песен
func GetSongs(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		group := r.URL.Query().Get("group")
		title := r.URL.Query().Get("title")
		page, _ := strconv.Atoi(r.URL.Query().Get("page"))
		limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))

		var songs []models.Song
		offset := (page - 1) * limit
		if err := db.Where("group LIKE ? AND title LIKE ?", "%"+group+"%", "%"+title+"%").Limit(limit).Offset(offset).Find(&songs).Error; err != nil {
			log.Printf("Error fetching songs: %v", err)
			http.Error(w, "Error fetching songs", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(songs)
	}
}

// CreateSong обрабатывает запрос на создание новой песни
func CreateSong(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var song models.Song
		if err := json.NewDecoder(r.Body).Decode(&song); err != nil {
			log.Printf("Error decoding request body: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Запрос к внешнему API для получения дополнительной информации о песне
		songDetail, err := services.GetSongDetails(song.Group, song.Title)
		if err != nil {
			log.Printf("Error fetching song details from external API: %v", err)
			http.Error(w, "Error fetching song details", http.StatusInternalServerError)
			return
		}

		// Обогащаем песню дополнительными данными
		song.ReleaseDate = songDetail.ReleaseDate
		song.Text = songDetail.Text
		song.Link = songDetail.Link

		// Сохраняем песню в базе данных
		db.Create(&song)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(song)
	}
}
