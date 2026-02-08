package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	http.Handle("/", templ.Handler(Hello("Rishav")))

	http.HandleFunc("/clicked", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<span>HTMX Swap Successful!</span>"))
	})

	fmt.Println("Server listening on :8085")
	http.ListenAndServe(":8085", nil)
}
