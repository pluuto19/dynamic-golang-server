package main

import (
	"html/template"
	"net/http"
	"log"
)

// Define a struct to pass dynamic data to templates
type PageData struct {
	Title string
	Content string
}

func main() {
	// Serve static files (CSS, images, etc.)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Define routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)

	// Start the server
	port := ":8080"
	log.Printf("Server running on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

// Home page handler
func homeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html", PageData{
		Title: "Welcome to My Go Website",
		Content: "This is the home page.",
	})
}

// About page handler
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html", PageData{
		Title: "About Us",
		Content: "This page tells you more about us.",
	})
}

// Helper function to render templates
func renderTemplate(w http.ResponseWriter, tmpl string, data PageData) {
	t, err := template.ParseFiles("templates/" + tmpl)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Println("Template Error:", err)
		return
	}
	t.Execute(w, data)
}
