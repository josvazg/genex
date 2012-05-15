package genrepo

import (
	"appengine"
	"appengine/datastore"
    "appengine/user"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

var templates = template.Must(template.ParseFiles("genrepo/browse.html" /*"genrepo/snippets.html"*/))

func init() {
    http.HandleFunc("/login", login)
    http.HandleFunc("/", browse)
}

const (
	PrimitivesOnly = iota // Means that the generic code is only applicable for primitive types or derivates
	Any                   // Any type is suitable to expand the generic code, primitive operators are not used
)

type GenericSnippet struct {
    Name, Desc string
    GenType    int
    Code       string
    Date       time.Time
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
    fmt.Println("u=", u)

    gs := &GenericSnippet{"a", "-", Any, "nothing yet", time.Now()}
    fmt.Println("gs=", gs)
    _, err := datastore.Put(c, datastore.NewIncompleteKey(c, "GenericSnippet", nil), gs)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    fmt.Println("*gs=", gs)

    q := datastore.NewQuery("GenericSnippet").Order("-Name").Limit(10)
    gss := make([]GenericSnippet, 0, 10)
    if _, err := q.GetAll(c, &gss); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    fmt.Println("gss=", gss)

    err = templates.ExecuteTemplate(w, "browse.html", u)
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

    fmt.Println("url=", url)
    w.Header().Set("Location", url)
	w.WriteHeader(http.StatusFound)
}

/*
func LoggedUser() string {
	return "-"
}*/
