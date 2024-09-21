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
            color: #e11d47;
        }
        .footer {
            font-size: 10pt;
        }
        .description {
            font-size: 14pt;
            margin-top: 20px;
            color: #9ca3af;
        }
        .current-url {
            font-size: 14pt;
            margin-top: 10px;
            color: #9ca3af;
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
    {{end}}
    {{if .Picked}}
    <p>Picked: {{.Picked}}</p>
    {{end}}
    {{if .ShowReload}}
    <p>Reload for different results!</p>
    {{end}}
    {{if .CurrentURL}}
    <p class="current-url">Current URL: <a href="{{.CurrentURL}}">{{.CurrentURL}}</a></p>
    {{end}}
    {{if .Description}}
    <p class="description">{{.Description}}</p>
    {{end}}
    <p class="footer">&copy; 2024 Klaus Breyer - <a href="https://www.v01.io/">v01.io</a></p>
</body>
</html>
`

const rootTmpl = `
<!DOCTYPE html>
<html>
<head>
    <title>Randomized Examples</title>
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
            color: #e11d47;
        }
        .footer {
            font-size: 10pt;
        }
        .description {
            font-size: 14pt;
            margin-top: 20px;
            color: #9ca3af;
        }
    </style>
</head>
<body>
    <h1>Randomized Examples</h1>
	<p class="description">This app generates a random list of names or selects a single name based on the specified route and displays them in HTML. It’s ideal for workshops, demonstrations, random selections in group activities, and rotating daily tasks where everyone sees the same result for the day.</p>
    <ul>
        <li><a href="/shuffle/klaus,linus,jonas,julia">Shuffle names: /shuffle/klaus,linus,jonas,julia</a></li>
        <p class="description">{{.ShuffleDescription}}</p>

        <li><a href="/shuffle-today/klaus,linus,jonas,julia">Shuffle today: /shuffle-today/klaus,linus,jonas,julia</a></li>
        <p class="description">{{.ShuffleTodayDescription}}</p>

        <li><a href="/pick/klaus,linus,jonas,julia">Pick a name: /pick/klaus,linus,jonas,julia</a></li>
        <p class="description">{{.PickDescription}}</p>

        <li><a href="/pick-today/klaus,linus,jonas,julia">Pick today: /pick-today/klaus,linus,jonas,julia</a></li>
        <p class="description">{{.PickTodayDescription}}</p>

    </ul>
    <p class="footer">&copy; 2024 Klaus Breyer - <a href="https://www.v01.io/">v01.io</a></p>
</body>
</html>
`

const (
	shuffleDescription      = "This route shuffles the given names randomly each time the page is reloaded."
	pickDescription         = "This route randomly picks a single name from the provided list on each reload."
	pickTodayDescription    = "This route picks a single name using today's date as the seed, resulting in the same name for the day."
	shuffleTodayDescription = "This route shuffles the list of names using today's date as the seed, producing a fixed order for the day."
)

type PageData struct {
	Names                   []string
	CurrentURL              string
	Picked                  string
	ShowReload              bool
	Description             string
	ShuffleDescription      string
	PickDescription         string
	PickTodayDescription    string
	ShuffleTodayDescription string
}

// randomizeNames handles the /shuffle/ route
func randomizeNames(w http.ResponseWriter, r *http.Request) {
	baseURL := getBaseURL(r)
	names := r.URL.Path[len("/shuffle/"):]
	parts := strings.FieldsFunc(names, func(c rune) bool {
		return c == ',' || c == ';'
	})

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(parts), func(i, j int) {
		parts[i], parts[j] = parts[j], parts[i]
	})

	data := PageData{
		Names:       parts,
		CurrentURL:  fmt.Sprintf("%s/shuffle/%s", baseURL, names),
		ShowReload:  true,
		Description: shuffleDescription,
	}
	renderTemplate(w, data)
}

// pickName handles the /pick/ route
func pickName(w http.ResponseWriter, r *http.Request) {
	baseURL := getBaseURL(r)
	names := r.URL.Path[len("/pick/"):]
	parts := strings.FieldsFunc(names, func(c rune) bool {
		return c == ',' || c == ';'
	})

	rand.Seed(time.Now().UnixNano())
	picked := parts[rand.Intn(len(parts))]

	data := PageData{
		Picked:      picked,
		CurrentURL:  fmt.Sprintf("%s/pick/%s", baseURL, names),
		ShowReload:  true,
		Description: pickDescription,
	}
	renderTemplate(w, data)
}

// pickNameToday handles the /pick-today/ route
func pickNameToday(w http.ResponseWriter, r *http.Request) {
	baseURL := getBaseURL(r)
	names := r.URL.Path[len("/pick-today/"):]
	parts := strings.FieldsFunc(names, func(c rune) bool {
		return c == ',' || c == ';'
	})

	// Seed with today's date
	seed := time.Now().Truncate(24 * time.Hour).Unix()
	rand.Seed(seed)
	picked := parts[rand.Intn(len(parts))]

	data := PageData{
		Picked:      picked,
		CurrentURL:  fmt.Sprintf("%s/pick-today/%s", baseURL, names),
		Description: pickTodayDescription,
	}
	renderTemplate(w, data)
}

// shuffleNamesToday handles the /shuffle-today/ route
func shuffleNamesToday(w http.ResponseWriter, r *http.Request) {
	baseURL := getBaseURL(r)
	names := r.URL.Path[len("/shuffle-today/"):]
	parts := strings.FieldsFunc(names, func(c rune) bool {
		return c == ',' || c == ';'
	})

	// Seed with today's date
	seed := time.Now().Truncate(24 * time.Hour).Unix()
	rand.Seed(seed)
	rand.Shuffle(len(parts), func(i, j int) {
		parts[i], parts[j] = parts[j], parts[i]
	})

	data := PageData{
		Names:       parts,
		CurrentURL:  fmt.Sprintf("%s/shuffle-today/%s", baseURL, names),
		Description: shuffleTodayDescription,
	}
	renderTemplate(w, data)
}

// rootHandler displays the main page with examples and descriptions
func rootHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		ShuffleDescription:      shuffleDescription,
		PickDescription:         pickDescription,
		PickTodayDescription:    pickTodayDescription,
		ShuffleTodayDescription: shuffleTodayDescription,
	}
	t := template.Must(template.New("root").Parse(rootTmpl))
	err := t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// renderTemplate renders the HTML template with the provided data
func renderTemplate(w http.ResponseWriter, data PageData) {
	t := template.Must(template.New("randomize").Parse(tmpl))
	err := t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// getBaseURL constructs the base URL dynamically from the request
func getBaseURL(r *http.Request) string {
	protocol := "http"
	if r.TLS != nil {
		protocol = "https"
	}
	return fmt.Sprintf("%s://%s", protocol, r.Host)
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/shuffle/", randomizeNames)
	http.HandleFunc("/pick/", pickName)
	http.HandleFunc("/pick-today/", pickNameToday)
	http.HandleFunc("/shuffle-today/", shuffleNamesToday)
	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
