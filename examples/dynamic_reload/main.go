package main

import (
	"net/http"

	"github.com/govenue/ace"
)

func handler(w http.ResponseWriter, r *http.Request) {
	tpl, err := ace.Load("example", "", &ace.Options{DynamicReload: true})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
