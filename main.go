package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	bind := os.Getenv("BIND")
	if bind == "" {
		bind = ":8080"
	}

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		param := query.Get("param")

		fmt.Fprintf(w, "param=%s", param)
		log.Printf("Receive GET request path = %s\n", r.RequestURI)
	})

	http.ListenAndServe(bind, nil)
}
