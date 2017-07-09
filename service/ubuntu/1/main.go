package main

import (
	"flag"
	"fmt"
	"net/http"
)

var addrs = flag.String("http", "localhost:65000", "Listen address")

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
	http.ListenAndServe(*addrs, nil)
}
