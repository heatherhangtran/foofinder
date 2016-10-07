package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I like %s!", r.URL.Path[1:])

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":5000", nil)
}
