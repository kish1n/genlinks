package data

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func getLink(db *sql.DB, shortened string) (string, error) {
	var original string

	err := db.QueryRow("SELECT original FROM links WHERE shortened = $1", shortened).Scan(&original)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil // Сокращенная ссылка не найдена
		}
		return "", fmt.Errorf("error checking existing link: %v", err)
	}

	return original, nil
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortened := vars["shortened"]

	db, err := initDB()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		log.Printf("RedirectHandler: %v", err)
		return
	}
	defer db.Close()

	original, err := getLink(db, shortened)
	if err != nil {
		http.Error(w, "Error getting link", http.StatusInternalServerError)
		log.Printf("GetLink: %v", err)
		return
	}

	if original == "" {
		http.Error(w, "Error not found links", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, original, http.StatusFound)
	log.Printf("Redirecting to: %s", original)
}
