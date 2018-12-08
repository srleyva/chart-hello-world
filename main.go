package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", sayHello)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from %s", r.URL.Path)
}
