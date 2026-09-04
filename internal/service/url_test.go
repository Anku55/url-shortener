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

func TestGetOriginalURL(t *testing.T) {
	store := storage.NewMemoryStorage()
	service := NewURLService(store)

	shortCode := "abc123"
	originalURL := "http://github.com"

	store.Set(shortCode, originalURL)

	returnedURL, found := service.GetOriginalURL(shortCode)

	if !found {
		t.Fatal("expected shortened URL to be found")
	}

	if returnedURL != originalURL {
		t.Fatalf("expected %s, got %s", originalURL, returnedURL)
	}
}

func TestGetOriginalURLNotFound(t *testing.T) {
	store := storage.NewMemoryStorage()
	service := NewURLService(store)

	shortCode := "doesNotExist"

	returnedURL, found := service.GetOriginalURL(shortCode)

	if found {
		t.Fatal("expected short code not to be found")
	}

	if returnedURL != "" {
		t.Fatalf("expected empty URL, got %s", returnedURL)
	}
}
