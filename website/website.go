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

// handler: basic handler for the web server
func handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("website.html")
	t.Execute(w, nil)
}

// homepage: tools page of the website
func homepage(w http.ResponseWriter, r *http.Request) {
	homepageNav := []Nav{
		{Href: "/", Name: "Homepage"},
		{Href: "/tools", Name: "Tools"},
		{Href: "/why", Name: "Why"},
	}
	t, _ := template.ParseFiles("homepage.html")
	t.Execute(w, homepageNav)
}

// tools: tools page of the website
func tools(w http.ResponseWriter, r *http.Request) {
	toolsNav := []Nav{
		{Href: "/", Name: "Homepage"},
		{Href: "/tools", Name: "Tools"},
		{Href: "/why", Name: "Why"},
	}
	t, _ := template.ParseFiles("tools.html")
	t.Execute(w, toolsNav)
}

// why: why page of the website
func why(w http.ResponseWriter, r *http.Request) {
	whyNav := []Nav{
		{Href: "/", Name: "Homepage"},
		{Href: "/tools", Name: "Tools"},
		{Href: "/why", Name: "Why"},
	}
	t, _ := template.ParseFiles("why.html")
	t.Execute(w, whyNav)
}

func main() {
	// convert port to string
	port := strconv.Itoa(PORT)

	// running a file
	fmt.Printf("Server running: http://localhost:%s\n", port)

	// web server starts at this directory
	fs := http.FileServer(http.Dir("."))

	http.Handle("/assets/", http.StripPrefix("/assets", fs))

	http.HandleFunc("/", handler)
	http.HandleFunc("/homepage", homepage)
	http.HandleFunc("/tools", tools)
	http.HandleFunc("/why", why)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
