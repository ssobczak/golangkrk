// Copyright 2013 Szymon Sobczak: http://about.me/ssobczak
// Licensed under the MIT license: http://opensource.org/licenses/MIT
// The above copyright notice shall be included in all copies or substantial portions of the Software.

package main

import (
	"fmt"
	"time"
)

type Ball struct{}

func player(number int, left_table chan Ball, right_table chan Ball) {
	ball := <-left_table
	// fmt.Printf("Player %d\n", number)
	right_table <- ball
}

const TABLES = 10e5

func main() {
	t := time.Now()

	tables := make([]chan Ball, TABLES)
	tables[0] = make(chan Ball)

	for i := 1; i != TABLES; i++ {
		tables[i] = make(chan Ball)
		go player(i, tables[i-1], tables[i])
	}

	fmt.Println("Let the game begin!")
	tables[0] <- Ball{}

	<-tables[TABLES-1]
	fmt.Println("Game over!")

	fmt.Printf("It took %v.\n", time.Since(t))
}
