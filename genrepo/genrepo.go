package genrepo

import (
	"appengine"
	"appengine/user"
	"fmt"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("genrepo/browse.html"))

func init() {
	http.HandleFunc("/login",login)
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
	c := appengine.NewContext(r)
    u := user.Current(c)
    fmt.Println("u=",u)
	err := templates.ExecuteTemplate(w, "browse.html", u)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func login(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	url, err := user.LoginURL(c, "/")
    if err != nil {
      	http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    fmt.Println("url=",url)
    w.Header().Set("Location", url)
	w.WriteHeader(http.StatusFound)
}
/*
func LoggedUser() string {
	return "-"
}*/
