package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
)

func getHTTPPort() string {
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		return "3000"
	}
	return port
}

func main() {
	/*
		TODO: csv2json
		GIVEN csv2json end point
		WHEN user give csv
		THEN json is returned
	*/

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", getHTTPPort()), nil))

}
