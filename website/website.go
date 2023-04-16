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

// handler: basic handler for the web server
func handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("website.html")
	t.Execute(w, nil)
}

func main() {
	// convert port to string
	port := strconv.Itoa(PORT)

	// running a file
	fmt.Printf("Server running: http://localhost:%s\n", port)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
