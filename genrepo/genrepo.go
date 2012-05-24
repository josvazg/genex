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

var templates = template.Must(template.ParseFiles("html/browse.html", "html/add.html",
    "html/templates.html","html/style.css"))

func init() {
    http.HandleFunc("/login", login)
    http.HandleFunc("/add",add)
    http.HandleFunc("/", browse)
}

const (
	PrimitivesOnly = iota // Means that the generic code is only applicable for primitive types or derivates
	Any                   // Any type is suitable to expand the generic code, primitive operators are not used
)

type GenSnippet struct {
    Name, Desc string
    GenType    int
    Code       string
    Date       time.Time
}

type Genex interface {
	list() []GenSnippet
	get(name string) GenSnippet
	put(gen GenSnippet)
	remove(name string)
}

type BrowsePage struct {
    User *user.User
    List []GenSnippet
}

func browse(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
    pg := &BrowsePage{user.Current(c), make([]GenSnippet, 0, 10)}

    q := datastore.NewQuery("GenSnippet").Order("-Name").Limit(10)
    if _, err := q.GetAll(c, &pg.List); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    fmt.Println("pg.List=", pg.List)

    err := templates.ExecuteTemplate(w, "browse.html", pg)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func add(w http.ResponseWriter, r *http.Request) {
    //c := appengine.NewContext(r)

    err := templates.ExecuteTemplate(w, "add.html", nil)
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
