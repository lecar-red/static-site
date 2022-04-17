package main

import (
	"log"
	"net/http"
)

func main() {
	port := ":8181"
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	log.Println("listening on ", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
