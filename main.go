package main

import (
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	http.Handle("/", templ.Handler(Hello("Rishav")))

	http.HandleFunc("/clicked", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<span>HTMX Swap Successful!</span>"))
	})

	http.ListenAndServe(":8080", nil)
}
