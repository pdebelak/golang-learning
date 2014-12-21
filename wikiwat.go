package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"net/http"
	"log"
	"os"
)


var db *sql.DB

func main() {
	var err error
	// this is so it works both locally and on heroku
	dbName := os.Getenv("DATABASE_URL")
	if dbName == "" {
		dbName = "dbname='wikiwat' sslmode='disable'"
	}
	db, err = sql.Open("postgres", dbName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))

	http.HandleFunc("/", staticHandler)
	http.HandleFunc("/page/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	// ditto
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)
}