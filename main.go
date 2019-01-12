/* What to do in each page:
 * "Home" should just show the intro to the reelo system.
 * "Ranks" should fetch the reelo data from the db (if requested re-run the reelo algorithm)
 * "Upload" should recive a ranking file and a format as input and recalculate the reelo socre. (should call a /save/)
 * " " should redirect to Home
 */
package main

import (
	"html/template"
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
// TODO this is useless atm
func loadPage(title string) (*Page, error) {
	body := []byte("")
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

func homeHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage("home")
	if err != nil {
		return
	}
	renderTemplate(w, "home", p)
}

func ranksHandler(w http.ResponseWriter, r *http.Request, title string) {
	// TODO: load ranking
	// select Nome, Cognome, reelo from Giocatore
	// inseriscili nella pagina
	p, err := loadPage(title)
	if err != nil {
		return
	}
	renderTemplate(w, title, p)
}

// TODO: maybe uploading is useless, we can do that from the server terminal
// the web service doesn't really need to be used by anyone.. maybe
func uploadHandler(w http.ResponseWriter, r *http.Request, title string) {
	// TODO: login?
	p, err := loadPage(title)
	if err != nil {
		return
	}
	renderTemplate(w, title, p)
}

func main() {
	http.HandleFunc("/home", makeHandler(homeHandler))
	http.HandleFunc("/ranks", makeHandler(ranksHandler))
	http.HandleFunc("/upload", makeHandler(uploadHandler))
	http.HandleFunc("/", makeHandler(homeHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
