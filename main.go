package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello You Have Made it to the server, %s \n", r.Host)
	})

	log.Fatal(http.ListenAndServe(":3030", nil))
}
