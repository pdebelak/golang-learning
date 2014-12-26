package main

import (
	"time"
	"log"
)

type Page struct {
	Id int64
	Title string `sql:"size:255"`
	Body string
	CreatedAt time.Time
	UpdatedAt time.Time
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

func recentPages() ([]string) {
	var titles []string
	var pages []Page
	err := db.Order("updated_at desc").Limit(5).Find(&pages).Pluck("title", &titles).Error
	if err != nil {
		log.Fatal(err)
	}
	return titles
}

func randomTitle() string {
	var page Page
	err := db.Order("RANDOM()").First(&page).Error
	if err != nil {
		log.Fatal(err)
	}
	return page.Title
}