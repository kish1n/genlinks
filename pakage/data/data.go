package data

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

func initDB() (*sql.DB, error) {
	connStr := fmt.Sprintf("%s:%s@%s:%s:%s?sslmode=disable",
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)

	fmt.Println("connStr: ", connStr)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	log.Println("Connected to the database successfully")

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS links (
		original TEXT NOT NULL,
		shortened TEXT PRIMARY KEY
	);`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		return nil, fmt.Errorf("error creating table: %v", err)
	}

	log.Println("Table 'links' checked/created successfully")

	return db, nil
}

func DBHandler(w http.ResponseWriter, r *http.Request) {
	db, err := initDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("DBHandler: %v", err)
		return
	}
	defer db.Close()

	var currentTime string
	err = db.QueryRow("SELECT NOW()").Scan(&currentTime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("DBHandler: %v", err)
		return
	}

	fmt.Fprintf(w, "Current time: %s", currentTime)
}
