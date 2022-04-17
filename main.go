package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	port := ":8181"
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", serveTemplate)

	log.Println("listening on ", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	log.Println("serving page:", r.URL.Path)

	lp := filepath.Join("templates", "layout.html")
	fp := filepath.Join("templates", filepath.Clean(r.URL.Path))

	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		//log.Fatal(err)
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	tmpl.ExecuteTemplate(w, "layout", nil)
}
