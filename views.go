package main

import (
	"html/template"
	"regexp"
	"github.com/russross/blackfriday"
	"github.com/microcosm-cc/bluemonday"
	"net/http"
	"log"
)

type Context struct {
	PageTitle string
	ThisPage *Page
	DisplayBody template.HTML
	List []string
}

var templates = template.Must(
				template.ParseFiles("tmpl/edit.html", 
								  	"tmpl/view.html", 
									"tmpl/home.html", 
									"tmpl/about.html",
									"tmpl/base.html"))
var validPath = regexp.MustCompile("^/(edit|save|page)/([a-zA-Z0-9]+)$")

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	unsafe := blackfriday.MarkdownCommon([]byte(p.Body))
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	var pt string
	if tmpl == "edit" {
		pt = "Editing "+p.Title
	} else {
		pt = p.Title
	}
	context := &Context{PageTitle: pt, ThisPage: p, DisplayBody: template.HTML(html)}
	err := templates.ExecuteTemplate(w, tmpl, context)
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
	http.Redirect(w, r, "/page/"+title, http.StatusFound)
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	p := template.HTMLEscapeString(r.URL.Path)
	if p == "/" {
		context := &Context{List: recentPages()}
		err := templates.ExecuteTemplate(w, "home", context)
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