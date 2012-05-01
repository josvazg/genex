package genrepo

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("genrepo/browse.html"))

func init() {
    http.HandleFunc("/", browse)
}

const (
	PrimitivesOnly = iota // Means that the generic code is only applicable for primitive types or derivates
	Any                   // Any type is suitable to expand the generic code, primitive operators are not used
)

type GenericSnippet struct {
	name, desc string
	genType int
	code string
}

type Genex interface {
	list() []GenericSnippet
	get(name string) GenericSnippet
	put(gen GenericSnippet)
	remove(name string)
}

func browse(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "browse.html", nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
