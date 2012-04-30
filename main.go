package main

import (
    "fmt"
    "net/http"
)

const (
	BrowsePage=`
<html>
<head>
<title>Genex: Go GENerics EXchange</title>
<style type="text/css">

body {
	font-family: Helvetica, Arial, sans-serif;
	margin: 0;
}

h1 {
	font-size: 18pt;
}

a {
	color: #375EAB;
	text-decoration: none;
}

b {
	color: #375EAB;
}

div.topbar {
	background-color: #E0EBF5;
	padding: 0.5em;
}

div.intro {
	margin: 2em;
}
</style>
</head>
<body>
<div class="topbar">
<h1>Genex: Go <b>Gen</b>erics <b>Ex</b>change</h1>
</div>
<div class="intro">
As of now (<a href="http://golang.org">Go 1</a>) does not support Generics within your Go code.<p/>

There are <a href="http://research.swtch.com/generic">very good reasons</a> for that, 
but while a solution is found to this dilema, programmers could use a little help avoiding generic code 
workarrounds that end up in repetitive and error prone coding tasks or lower performance.<p/>

Genex is a Go Generics Exchange developer helper tool. It has two components:<br/>
<ol>
<li value="1">This server side generics code snippets repository.</li>
<li>A command line tool to download snippets from expanded with your favorite types.</li>
</ol>

</div>
</body>
</html>
`
)

func browse(w http.ResponseWriter, r *http.Request) {
    //fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
    fmt.Fprintf(w,BrowsePage)
}

func main() {
    http.HandleFunc("/", browse)
    http.ListenAndServe(":8080", nil)
}