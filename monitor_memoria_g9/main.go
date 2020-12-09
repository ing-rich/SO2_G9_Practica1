package main

import (
	"net/http"
)

func main() {
	// routes
	http.HandleFunc("/", homeHandler)
	//start the server
	http.ListenAndServe(":3000", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
