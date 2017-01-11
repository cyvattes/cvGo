package main

import (
	"net/http"
)

func pageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func main() {
	http.HandleFunc("/", pageHandler)
	http.ListenAndServe(":8080", nil)
}
