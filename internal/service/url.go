package service

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net/url"
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
	err:=validateURL(originalURL)
	if err!=nil{
		return "",err
	}
	shortCode, err := generateShortCode(6)
	if err != nil {
		return "", err
	}

	s.storage.Set(shortCode, originalURL)

	return shortCode, nil
}
func (s *URLService) GetOriginalURL(shortCode string) (string, bool) {
	return s.storage.Get(shortCode)
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

func validateURL(originalURL string) error {
	parsedURL, err := url.Parse(originalURL)

	if err != nil {
		return err
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return fmt.Errorf("invalid scheme: %s (only http and https are allowed)", parsedURL.Scheme)

	}

	if parsedURL.Host == "" {
		return fmt.Errorf("url host cannot be empty")
	}


	return nil
}
