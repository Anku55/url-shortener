package main

import (
	"fmt"
	"net/http"

	"github.com/Anku55/url-shortener/internal/handler"
	"github.com/Anku55/url-shortener/internal/service"
	"github.com/Anku55/url-shortener/internal/storage"
)

func main() {
	store := storage.NewMemoryStorage()

	urlService := service.NewURLService(store)

	urlHandler := handler.NewURLHandler(urlService)

	http.HandleFunc("/shorten", urlHandler.Shorten)
	http.HandleFunc("/", urlHandler.Redirect)

	fmt.Println("Server running on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}