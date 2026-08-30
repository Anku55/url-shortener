package storage

import "testing"

func TestMemoryStorage_SetAndGet(t *testing.T) {
	storage := NewMemoryStorage()

	storage.Set("abc123", "https://github.com")

	url, exists := storage.Get("abc123")

	if !exists {
		t.Fatal("expected URL to exist")
	}

	if url != "https://github.com" {
		t.Fatalf("expected https://github.com, got %s", url)
	}
}