package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

const tmpl = `
<!DOCTYPE html>
<html>
<head>
    <title>Randomized</title>
    <style>
        body {
            background-color: #1e1e1e;
            color: #d4d4d4;
            font-size: 20pt;
            font-family: -apple-system, BlinkMacSystemFont, avenir next, avenir, segoe ui, helvetica neue, helvetica, Cantarell, Ubuntu, roboto, noto, arial, sans-serif;
        }
        a, a:visited, a:hover, a:active {
            color: #d4d4d4;
        }
        h1 {
            color: #db2777;
        }
        .footer {
            font-size: 10pt;
        }
    </style>
</head>
<body>
    <h1>Randomized</h1>
    {{if .Names}}
    <ul>
    {{range .Names}}
        <li>{{.}}</li>
    {{end}}
    </ul>
    {{else}}
    <p>No names provided. Try it with a sample set of names:
    <a href="/steve,tim,johnny">steve,tim,johnny</a></p>
    {{end}}
	{{if .Names}}
	<p>Reload for different results!</p>
    <p>Current URL: <a href="{{.CurrentURL}}">{{.CurrentURL}}</a></p>
	{{end}}

    <p class="footer">&copy; 2024 Klaus Breyer - <a href="https://www.v01.io/">v01.io</a></p>
</body>
</html>
`

type PageData struct {
	Names      []string
	CurrentURL string
}

func randomizeNames(w http.ResponseWriter, r *http.Request) {
	names := r.URL.Path[len("/"):]

	if names == "" {
		data := PageData{
			CurrentURL: "http://localhost:8080",
		}
		t := template.Must(template.New("randomize").Parse(tmpl))
		err := t.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	parts := strings.FieldsFunc(names, func(c rune) bool {
		return c == ',' || c == ';'
	})

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(parts), func(i, j int) {
		parts[i], parts[j] = parts[j], parts[i]
	})

	data := PageData{
		Names:      parts,
		CurrentURL: fmt.Sprintf("http://localhost:8080/%s", names),
	}

	t := template.Must(template.New("randomize").Parse(tmpl))
	err := t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", randomizeNames)
	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
