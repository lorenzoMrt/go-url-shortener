package controllers

import (
	"database/sql"
	"github.com/lorenzoMrt/go-url-shortener/internal/db"
	"github.com/lorenzoMrt/go-url-shortener/internal/url"
	"html/template"
	"net/http"
	"strings"
)

func Shorten(lite *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		if err := db.StoreUrl(lite, shortUrl, originalUrl); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

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
}

func Redirect(lite *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shortUrl := r.URL.Path[1:]
		if shortUrl == "" {
			http.Error(w, "Url required", http.StatusBadRequest)
			return
		}
		originUrl, err := db.GetOriginalUrl(lite, shortUrl)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Redirect(w, r, originUrl, http.StatusPermanentRedirect)
	}
}
