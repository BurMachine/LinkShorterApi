package interfaces

// Storage - Интерфейс хранилища
type Storage interface {
	GenerateShortLink()
	GetFullLink()
}
