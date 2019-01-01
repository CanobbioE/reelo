package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

const HTML_PATH = "./html/"

type Page struct {
	Title string
	Body  []byte
}

// loadPage loads a the page with the given title
func loadPage(title string) (*Page, error) {
	filename := HTML_PATH + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// renderTemplate draws the html page
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(HTML_PATH + tmpl + ".html")
	t.Execute(w, p)
}

var validPath = regexp.MustCompile("([a-zA-Z0-9]+)")

// makeHandler returns a http Handler Func
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

// TODO
func todo(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		return
	}
	renderTemplate(w, "todo", p)
}

func main() {
	// TODO: serve all the pages
	http.HandleFunc("/todo/", makeHandler(todo))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
