// Copyright 2013 Szymon Sobczak: http://about.me/ssobczak
// Licensed under the MIT license: http://opensource.org/licenses/MIT
// The above copyright notice shall be included in all copies or substantial portions of the Software.

package main

import (
	"encoding/json"
	"fmt"
	"github.com/deckarep/golang-set"
	"github.com/ssobczak/golangkrk/latency_workshop/distance"
	"log"
	"net/http"
	"runtime"
	"strings"
)

const (
	url = "localhost:1234"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/find", handler)
	http.HandleFunc("/stack", stack)

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

func stack(resp http.ResponseWriter, req *http.Request) {
	buf := make([]byte, 1e6)
	size := runtime.Stack(buf, true)
	fmt.Fprint(resp, string(buf[:size]))
}

type scored struct {
	Id       string
	Distance int
}

func get_distance_async(id, sequence string, responses chan scored) {
	responses <- scored{id, distance.GetDistance(id, sequence)}
}

func score(ids []string, sequence string) map[string]int {
	responses := make(chan scored)
	to_score := mapset.NewSet()

	for _, id := range ids {
		to_score.Add(id)
		go get_distance_async(id, sequence, responses)
	}

	result := make(map[string]int)
	for to_read := len(ids); to_read != 0; to_read-- {
		resp := <-responses
		result[resp.Id] = resp.Distance

		to_score.Remove(resp.Id)
		if to_score.Size() == len(ids)/5 {
			to_read += to_score.Size()
			go add_duplicates(to_score.Clone(), sequence, responses)
		}

		if to_score.Size() == 0 {
			go cleanup(to_read, responses)
			break
		}
	}

	return result
}

func add_duplicates(duplicated mapset.Set, sequence string, responses chan scored) {
	for id, _ := range duplicated {
		str_id, _ := id.(string)
		go get_distance_async(str_id, sequence, responses)
	}
}

func cleanup(to_read int, responses chan scored) {
	for i := 0; i != to_read-1; i++ {
		<-responses
	}
	close(responses)
}
