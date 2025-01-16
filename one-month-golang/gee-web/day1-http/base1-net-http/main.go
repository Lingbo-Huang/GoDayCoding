package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)      // curl http://localhost:9999/
	http.HandleFunc("/hello", helloHandler) // curl http://localhost:9999/hello
	log.Fatal(http.ListenAndServe(":9999", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path= %q\n", r.URL.Path)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
