package cli

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kish1n/genlinks/internal/data"
)

func Run(args []string) bool {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the URL shortener!"))
	})
	r.Get("/db", data.DBHandler)
	r.Post("/add", data.AddLinkHandler)
	r.Get("/{shortened}", data.RedirectHandler)

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
		return false
	}
	return true
}
