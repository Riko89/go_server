package main

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func renderTemplate(w http.ResponseWriter, name string, data interface{}) {
	tmpl := template.Must(template.ParseFiles(
		filepath.Join("templates", "base.html"),
		filepath.Join("templates", name), // name is "home.html"
	))
	tmpl.ExecuteTemplate(w, "base", data) // must match the layout name	
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<p>Hello from Go!!</p>"))
}


func main() {
	http.HandleFunc("/", homeHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/hello", helloHandler)

	println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
