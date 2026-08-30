package service

import (
	"crypto/rand"
	"math/big"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type Storage interface {
	Set(shortCode string, originalURL string)
	Get(shortCode string) (string, bool)
}

type URLService struct {
	storage Storage
}

func NewURLService(storage Storage) *URLService {
	return &URLService{
		storage: storage,
	}
}

func (s *URLService) ShortenURL(originalURL string) (string, error) {
	shortCode, err := generateShortCode(6)
	if err != nil {
		return "", err
	}

	s.storage.Set(shortCode, originalURL)

	return shortCode, nil
}

func generateShortCode(length int) (string, error) {
	result := make([]byte, length)

	for i := range result {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}

		result[i] = charset[n.Int64()]
	}

	return string(result), nil
}