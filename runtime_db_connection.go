package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":9999", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hellow")
}
