/* What to do in each page:
 * "Home" should just show the intro to the reelo system.
 * "Ranks" should fetch the reelo data from the db (if requested re-run the reelo algorithm)
 * "Upload" should recive a ranking file and a format as input and recalculate the reelo socre. (should call a /save/)
 * " " should redirect to Home
 */
package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

const ROOT = "."
const HTML_PATH = ROOT + "/html/"
const LAYOUT = HTML_PATH + "layout.html"

type Page struct {
	Title string
	Body  []byte
}

// loadPage loads a the page with the given title
// TODO: Probably there's no need for the body to be saved in .txt
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
	t, _ := template.ParseFiles(HTML_PATH+tmpl+".html", LAYOUT)
	t.ExecuteTemplate(w, "layout", p)
}

// makeHandler returns a http Handler Func
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	validPath := regexp.MustCompile("^/(home|ranks|upload)?")
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[1])
	}
}

// TODO different handers
func handler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		return
	}
	renderTemplate(w, title, p)
}

func main() {
	http.HandleFunc("/home", makeHandler(handler))
	http.HandleFunc("/ranks", makeHandler(handler))
	http.HandleFunc("/upload", makeHandler(handler))
	http.HandleFunc("/", makeHandler(handler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
