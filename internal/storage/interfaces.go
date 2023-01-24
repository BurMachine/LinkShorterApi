package storage

// ServiceStorage - Интерфейс хранилища
type ServiceStorage interface {
	AddShortLink(link, shortLink string) error
	GetFullLink(shortLink string) (string, error)
}
