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

type Tool struct {
	Title string
	Cmd string
	Desc string
}

type ToolPage struct {
	Navs []Nav
	Tools []Tool
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
	tools := []Tool{
		{Title: "Help", Cmd: "help", Desc: "This command shows a list of commands or help for one command"},
		{Title: "Weather", Cmd: "wttr", Desc: "This command shows you the weather forecast for the next couple of days (with help of wttr.in)"},
		{Title: "Dance", Cmd: "dance", Desc: "See Ice Spice's most famous video on the terminal with ascii art"},
		{Title: "Logo", Cmd: "logo", Desc: "This command displays our logo from the homepage in ascii art"},
		{Title: "Text", Cmd: "text", Desc: "This command sends a text message from a Miami number to person of your choice"},
		{Title: "Name Servers", Cmd: "ns", Desc: "This command shows current name servers for url"},
		{Title: "IP Address", Cmd: "ip", Desc: "This command shows both the IPv4 and IPv6 address of url"},
		{Title: "Operating System", Cmd: "os", Desc: "This command checks the operating system & architecture of current device"},
		//{Title: "", Cmd: "", Desc: ""},		
	}

	toolpage := ToolPage{
		Navs: toolsNav,
		Tools: tools,
	}
	
	t.ExecuteTemplate(w, "tools", toolpage)
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
