// Copyright 2013 Szymon Sobczak: http://about.me/ssobczak
// Licensed under the MIT license: http://opensource.org/licenses/MIT
// The above copyright notice shall be included in all copies or substantial portions of the Software.

package distance

import (
	"bufio"
	"net/http"
)

func GetDistance(uniprot_id, what string) int {
	return distance(what, fetch(uniprot_id))
}

func fetch(uniprot_id string) string {
	resp, _ := http.Get("http://www.uniprot.org/uniprot/" + uniprot_id + ".fasta")
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	scanner.Scan() // discard comment line

	fasta := ""
	for scanner.Scan() {
		fasta += scanner.Text()
	}

	return fasta
}

func distance(what, where string) int {
	best := len(what)

	for i := 0; i != len(where)-len(what); i++ {
		score := 0
		for j := 0; j != len(what); j++ {
			if what[j] != where[i+j] {
				score++
			}
		}
		if score < best {
			best = score
		}
	}

	return best
}
