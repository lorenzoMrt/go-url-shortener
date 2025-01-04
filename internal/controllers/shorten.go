package controllers

import (
	"github.com/lorenzoMrt/go-url-shortener/internal/url"
	"html/template"
	"net/http"
	"strings"
)

func Shorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	originalUrl := r.FormValue("url")
	if originalUrl == "" {
		http.Error(w, "Url required", http.StatusBadRequest)
		return
	}

	if !strings.HasPrefix(originalUrl, "http://") || !strings.HasPrefix(originalUrl, "https://") {
		originalUrl = "https://" + originalUrl
	}

	shortUrl := url.Shorten(originalUrl)
	data := map[string]string{
		"ShortURL": shortUrl,
	}

	t, err := template.ParseFiles("internal/views/shorten.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
