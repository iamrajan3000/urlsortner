package store

type Store interface {
	SaveURL(url string, url2 string)
	GetOriginalURL(url string) (string, bool)
}
