package main

import (
	"html/template"
	"database/sql"
	_ "github.com/lib/pq"
	"net/http"
	"regexp"
	"log"
	"time"
	"os"
)

type Page struct {
	Title string
	Body string
}

var templates = template.Must(template.ParseFiles("tmpl/edit.html", 
												  "tmpl/view.html", 
												  "tmpl/home.html", 
												  "tmpl/about.html",
												  "tmpl/base.html"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
var db *sql.DB

func (p *Page) save() error {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	_, err = tx.Exec("update pages set body = $1, updated_at = $2 where title = $3", p.Body, time.Now(), p.Title)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
	return err
}

func (p *Page) create() error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("insert into pages (title, created_at) values ($1, $2)", p.Title, time.Now())
	if err != nil {
		return err
	}
	tx.Commit()
	return err
}

func loadPage(title string) (*Page, error) {
	stmt, err := db.Prepare("select body from pages where title = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var body string
	err = stmt.QueryRow(title).Scan(&body)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func makeHandler (fn func (http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
		err = p.create()
		if err != nil {
			log.Fatal(err)
		}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: body}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	p := template.HTMLEscapeString(r.URL.Path)
	if p == "/" {
		err := templates.ExecuteTemplate(w, "home", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else if p == "/about" {
		err := templates.ExecuteTemplate(w, "about", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

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

	http.HandleFunc("/", staticHandler)
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	// ditto
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)
}