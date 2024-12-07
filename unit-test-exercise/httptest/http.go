package httptest

import (
	"log"
	"net/http"
)

func Web(addr string) {
	http.HandleFunc("/hello", HelloHandler)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!"))
}
