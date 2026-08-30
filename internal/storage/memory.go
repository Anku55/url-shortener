package storage

type MemoryStorage struct {
	urls map[string]string
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		urls: make(map[string]string),
	}
}

func (m *MemoryStorage) Set(shortCode string, originalURL string) {
	m.urls[shortCode] = originalURL
}

func (m *MemoryStorage) Get(shortCode string) (string, bool) {
	value, exists := m.urls[shortCode]
	if !exists {
		return "", false
	}
	return value, true
}
