package pagestatus

import (
	"log"
	"net/http"
)

// Page contains the struct with the data from the page
type Page struct {
	PageURL    string
	StatusCode int
}

// CheckPage verify the status of a web page
func (p *Page) CheckPage() {
	if p.PageURL == "" {
		log.Println("Undefined page")
		return
	}

	r, e := http.Get(p.PageURL)
	if e != nil {
		log.Println("Can't get the page")
		p.StatusCode = -1
	}
	p.StatusCode = r.StatusCode
}
