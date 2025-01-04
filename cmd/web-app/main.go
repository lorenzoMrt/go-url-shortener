package main

import (
	"github.com/lorenzoMrt/go-url-shortener/internal/controllers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/shorten", controllers.Shorten)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
