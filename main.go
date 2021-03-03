package main

import (
	"github.com/TheEgid/news-demo-go/views"
	"net/http"
)

func main() {
	port := "80"

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("assets"))

	mux.HandleFunc("/search", views.SearchHandler)
	mux.HandleFunc("/", views.IndexHandler)

	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	_ = http.ListenAndServe(":"+port, mux)
}

//
