package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

const PORT = 8081

type WebPage struct {
	Homepage string
}

type Nav struct {
	Href string
	Name string
}

// homepage: tools page of the website
func homepage(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("*.html"))
	homepageNav := []Nav{
		{Href: "/", Name: "Homepage"},
		{Href: "/tools", Name: "Tools"},
		{Href: "/why", Name: "Why"},
	}
	t.ExecuteTemplate(w, "homepage", homepageNav)
}

// tools: tools page of the website
func tools(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("*.html"))	
	toolsNav := []Nav{
		{Href: "/", Name: "Homepage"},
		{Href: "/tools", Name: "Tools"},
		{Href: "/why", Name: "Why"},
	}
	t.ExecuteTemplate(w, "tools", toolsNav)
}

// why: why page of the website
func why(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("*.html"))
	whyNav := []Nav{
		{Href: "/", Name: "Homepage"},
		{Href: "/tools", Name: "Tools"},
		{Href: "/why", Name: "Why"},
	}
	t.ExecuteTemplate(w,"why", whyNav)
}

func main() {
	// convert port to string
	port := strconv.Itoa(PORT)

	// running a file
	fmt.Printf("Server running: http://localhost:%s\n", port)

	// web server starts at this directory
	fs := http.FileServer(http.Dir("."))

	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	http.HandleFunc("/", homepage)
	http.HandleFunc("/homepage", homepage)
	http.HandleFunc("/tools", tools)
	http.HandleFunc("/why", why)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
