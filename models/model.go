package models

import "sync"

// URLStore holds the mappings between original and shortened URLs
type URLStore struct {
	sync.RWMutex
	urlMap map[string]string
}

type RequestData struct {
	URL string `json:"url"`
}
