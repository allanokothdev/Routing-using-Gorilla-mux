package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// Create a new request router that will receive all requests and pass it to request handler
	r := mux.NewRouter()

	 r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        title := vars["title"]
        page := vars["page"]

        fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
    })

    http.ListenAndServe(":80", r)
}