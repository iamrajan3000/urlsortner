package services

type Servicer interface {
	ShortenURL(originalURL string) string
	GetOriginalURL(shortURL string) (string, bool)
}
