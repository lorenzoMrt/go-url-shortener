package main

import (
	"database/sql"
	"github.com/lorenzoMrt/go-url-shortener/internal/controllers"
	"github.com/lorenzoMrt/go-url-shortener/internal/db"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

func main() {
	sqlite, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer func(sqlite *sql.DB) {
		err := sqlite.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(sqlite)

	if err := db.CreateTable(sqlite); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			controllers.Home(writer, r)
		} else {
			controllers.Redirect(sqlite)(writer, r)
		}
	})
	http.HandleFunc("/shorten", controllers.Shorten(sqlite))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
