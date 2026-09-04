package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Anku55/url-shortener/internal/service"
	"github.com/Anku55/url-shortener/internal/storage"
)

func TestRedirect(t *testing.T) {
	store := storage.NewMemoryStorage()
	urlService := service.NewURLService(store)
	urlHandler := NewURLHandler(urlService)

	shortCode := "abc123"
	originalURL := "https://github.com"

	store.Set(shortCode, originalURL)

	req := httptest.NewRequest(http.MethodGet, "/abc123", nil)
	recorder := httptest.NewRecorder()

	urlHandler.Redirect(recorder, req)

	if recorder.Code != http.StatusFound {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusFound,
			recorder.Code,
		)
	}

	location := recorder.Header().Get("Location")

	if location != originalURL {
		t.Fatalf(
			"expected location %s, got %s",
			originalURL,
			location,
		)
	}
}

func TestRedirectNotFound(t *testing.T) {
	store := storage.NewMemoryStorage()
	urlService := service.NewURLService(store)
	urlHandler := NewURLHandler(urlService)

	req := httptest.NewRequest(http.MethodGet, "/doesNotExist", nil)
	recorder := httptest.NewRecorder()

	urlHandler.Redirect(recorder, req)

	if recorder.Code != http.StatusNotFound {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusNotFound,
			recorder.Code,
		)
	}
}

func TestRedirectMethodNotAllowed(t *testing.T) {
	store := storage.NewMemoryStorage()
	urlService := service.NewURLService(store)
	urlHandler := NewURLHandler(urlService)

	req := httptest.NewRequest(http.MethodPost, "/abc123", nil)
	recorder := httptest.NewRecorder()

	urlHandler.Redirect(recorder, req)

	if recorder.Code != http.StatusMethodNotAllowed {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusMethodNotAllowed,
			recorder.Code,
		)
	}
}

func TestShortenInvalidURL(t *testing.T) {
	store := storage.NewMemoryStorage()
	service := service.NewURLService(store)
	handler := NewURLHandler(service)

	body := `{"url":"hello"}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/shorten",
		strings.NewReader(body),
	)

	recorder := httptest.NewRecorder()

	handler.Shorten(recorder, req)

	response := recorder.Result()

	if response.StatusCode != http.StatusBadRequest {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusBadRequest,
			response.StatusCode,
		)
	}
}
