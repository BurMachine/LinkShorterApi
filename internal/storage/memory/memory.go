package memory

type InMemoryStorage struct {
	Storage map[string]string
}

func NewStorageInit() *InMemoryStorage {
	s := InMemoryStorage{Storage: make(map[string]string)}
	return &s
}

func (s *InMemoryStorage) GenerateShortLink() {
	//TODO implement me
	panic("implement me")
}

func (s *InMemoryStorage) GetFullLink() {
	//TODO implement me
	panic("implement me")
}
