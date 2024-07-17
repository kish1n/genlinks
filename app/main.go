package main

import (
	"genlinks/data"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the URL shortener!"))
	})
	r.HandleFunc("/db", data.DBHandler)
	r.HandleFunc("/add", data.AddLinkHandler).Methods("POST")
	r.HandleFunc("/{shortened}", data.RedirectHandler).Methods("GET")

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r)) // Передаем маршрутизатор r
}
