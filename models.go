package main

import (
	"time"
	"log"
)

type Page struct {
	Title string
	Body string
	CreatedAt time.Time
	UpdatedAt time.Time
}

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

func (p *Page) FormattedCreate() string {
	return formatTime(p.CreatedAt)
}

func (p *Page) FormattedUpdate() string {
	return formatTime(p.UpdatedAt)
}

func formatTime(aTime time.Time) string {
	return aTime.Format("Mon, Jan _2, 2006 at 15:04:05")	
}

func loadPage(title string) (*Page, error) {
	stmt, err := db.Prepare("select body, updated_at, created_at from pages where title = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var (
		body string
		updated_at time.Time
		created_at time.Time
	)
	err = stmt.QueryRow(title).Scan(&body, &updated_at, &created_at)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body, CreatedAt: created_at, UpdatedAt: updated_at}, nil
}

func recentPages() ([]string) {
	var pages []string
	var title string
	rows, err := db.Query("select title from pages order by updated_at desc limit 5")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&title)
		if err != nil {
			log.Fatal(err)
		}
		pages = append(pages, title)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return pages
}