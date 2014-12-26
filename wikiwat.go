package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"net/http"
	"log"
	"os"
)


var db gorm.DB

func main() {
	var err error
	// this is so it works both locally and on heroku
	dbName := os.Getenv("DATABASE_URL")
	if dbName == "" {
		dbName = "dbname='wikiwat' sslmode='disable'"
	}
	db, err = gorm.Open("postgres", dbName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.AutoMigrate(&Page{})
	db.Model(&Page{}).AddIndex("idx_page_title", "title")

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))

	http.HandleFunc("/", staticHandler)
	http.HandleFunc("/page/random", randomHandler)
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