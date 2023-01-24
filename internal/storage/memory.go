package storage

import "burmachine/LinkGenerator/internal/serviceErrors"

type InMemoryStorage struct {
	StorageShortKey    map[string]string
	StorageOriginalKey map[string]string
}

func NewInMemoryStorageInit() *InMemoryStorage {
	s := InMemoryStorage{StorageShortKey: make(map[string]string), StorageOriginalKey: make(map[string]string)}
	return &s
}

func (s *InMemoryStorage) AddShortLink(link, shortLink string) error {
	err := s.CheckLink(link, shortLink)
	if err != nil {
		return err
	}
	s.StorageOriginalKey[link] = shortLink
	s.StorageShortKey[shortLink] = link
	return nil
}

func (s InMemoryStorage) GetFullLink(shortLink string) (string, error) {

	originalLink, ok := s.StorageShortKey[shortLink]
	if !ok {
		return originalLink, serviceErrors.ErrorLinkNotExist
	}
	return s.StorageShortKey[shortLink], nil

}

func (s *InMemoryStorage) CheckLink(link, shortLink string) error {
	_, ok := s.StorageShortKey[shortLink]
	if ok {
		return serviceErrors.ErrorLinkExist
	}
	_, ok = s.StorageShortKey[link]
	if ok {
		return serviceErrors.ErrorLinkExist
	}
	return nil
}
