package service

import (
	"testing"

	"github.com/Anku55/url-shortener/internal/storage"
)

func TestShortenURL(t *testing.T) {
	store := storage.NewMemoryStorage()
	service := NewURLService(store)

	originalURL := "https://github.com"

	shortCode, err := service.ShortenURL(originalURL)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(shortCode) != 6 {
		t.Fatalf("expected short code length 6, got %d", len(shortCode))
	}

	url, exists := store.Get(shortCode)

	if !exists {
		t.Fatal("expected shortened URL to exist in storage")
	}

	if url != originalURL {
		t.Fatalf("expected %s, got %s", originalURL, url)
	}
}