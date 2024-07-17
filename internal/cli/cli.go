package cli

import (
	"github.com/gorilla/mux"
	"github.com/shortener/internal/data"
	"log"
	"net/http"
)

func Run(args []string) bool {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the URL shortener!"))
	})
	r.HandleFunc("/db", data.DBHandler)
	r.HandleFunc("/add", data.AddLinkHandler).Methods("POST")
	r.HandleFunc("/{shortened}", data.RedirectHandler).Methods("GET")

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
		return false
	}
	return true
}
