package store

import (
	"sync"
)

// store holds the mappings between original and shortened URLs
type store struct {
	sync.RWMutex
	URLMap map[string]string
}

// New initializes the store
func New() Store {
	return &store{
		URLMap: make(map[string]string),
	}
}

// SaveURL saves the shortened URL and original URL in the store
func (store *store) SaveURL(shortURL, originalURL string) {
	store.Lock()
	defer store.Unlock()
	store.URLMap[shortURL] = originalURL
}

// GetOriginalURL retrieves the original URL from the store using the shortened URL
func (store *store) GetOriginalURL(shortURL string) (string, bool) {
	store.RLock()
	defer store.RUnlock()
	originalURL, exists := store.URLMap[shortURL]
	return originalURL, exists
}
