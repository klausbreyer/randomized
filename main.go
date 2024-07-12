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
    <title>Parameter Roulette</title>
    <style>
        body {
            background-color: #1e1e1e;
            color: #d4d4d4;
            font-size: 20pt;
            font-family: -apple-system, BlinkMacSystemFont, avenir next, avenir, segoe ui, helvetica neue, helvetica, Cantarell, Ubuntu, roboto, noto, arial, sans-serif;
        }
        h1 {
            color: #db2777;
        }
    </style>
</head>
<body>
    <h1>Parameter Roulette</h1>
    <ul>
    {{range .}}
        <li>{{.}}</li>
    {{end}}
    </ul>
</body>
</html>
`

func randomizeNames(w http.ResponseWriter, r *http.Request) {
	names := r.URL.Path[len("/"):]

	parts := strings.FieldsFunc(names, func(c rune) bool {
		return c == ',' || c == ';'
	})

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(parts), func(i, j int) {
		parts[i], parts[j] = parts[j], parts[i]
	})

	t := template.Must(template.New("randomize").Parse(tmpl))
	err := t.Execute(w, parts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", randomizeNames)
	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
