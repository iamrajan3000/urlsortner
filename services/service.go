package services

import (
	"urlsortner/store"
	"urlsortner/utils"
)

// Service contains business logic related to URLs
type Service struct {
	Store store.Store
}

// New initializes the Service
func New(store store.Store) Servicer {
	return &Service{
		Store: store,
	}
}

// ShortenURL generates a short URL and saves it in the store
func (s *Service) ShortenURL(originalURL string) string {
	shortURL := utils.GenerateShortURL()
	s.Store.SaveURL(shortURL, originalURL)
	return shortURL
}

// GetOriginalURL fetches the original URL using the short URL
func (s *Service) GetOriginalURL(shortURL string) (string, bool) {
	return s.Store.GetOriginalURL(shortURL)
}
