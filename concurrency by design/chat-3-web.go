package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello Web!")
	})

	if err := http.ListenAndServe("localhost:1234", nil); err != nil {
		log.Fatal(err)
	}
}
