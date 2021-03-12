package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
)

// appContent is our static web server content.
//go:embed app/*
var appContent embed.FS

// tplContent is our templates
//go:embed tpl/*
var tplContent embed.FS

func main() {
	mux := http.NewServeMux()
	mux.Handle("/app/", http.FileServer(http.FS(appContent)))
	mux.HandleFunc("/v1/values", valuesHandler)
	template.ParseFS(tplContent, "tpl/*.tmpl")
	panic(http.ListenAndServe(":7000", mux))
}

func valuesHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("handle %s", r.URL.Path)
	http.Error(w, "not implemented", http.StatusInternalServerError)
}
