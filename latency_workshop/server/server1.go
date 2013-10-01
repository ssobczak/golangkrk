// Copyright 2013 Szymon Sobczak: http://about.me/ssobczak
// Licensed under the MIT license: http://opensource.org/licenses/MIT
// The above copyright notice shall be included in all copies or substantial portions of the Software.

package main

import (
	"encoding/json"
	"fmt"
	"github.com/ssobczak/golangkrk/latency_workshop/distance"
	"log"
	"net/http"
	"strings"
)

const (
	url = "localhost:1234"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/find", handler)

	fmt.Println("Listenning on " + url)
	if err := http.ListenAndServe(url, nil); err != nil {
		log.Fatal(err)
	}
}

func handler(resp http.ResponseWriter, req *http.Request) {
	ids := strings.Split(req.PostFormValue("ids"), "\n")
	sequence := req.PostFormValue("sequence")

	serialized, _ := json.Marshal(score(ids, sequence))
	fmt.Fprint(resp, string(serialized))
}

type scored struct {
	Id       string
	Distance int
}

func score(ids []string, sequence string) map[string]int {
	responses := make(chan scored)
	defer close(responses)

	for _, id := range ids {
		go func(id string) {
			responses <- scored{id, distance.GetDistance(id, sequence)}
		}(id)
	}

	result := make(map[string]int)
	for _ = range ids {
		resp := <-responses
		result[resp.Id] = resp.Distance
	}

	return result
}
