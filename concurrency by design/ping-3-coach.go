// Copyright 2013 Szymon Sobczak: http://about.me/ssobczak
// Licensed under the MIT license: http://opensource.org/licenses/MIT
// The above copyright notice shall be included in all copies or substantial portions of the Software.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Ball struct{}

func coach(player chan string) {
	for {
		switch rand.Int() % 3 {
		case 0:
			player <- "forehand"
		case 1:
			player <- "backhand"
		case 2:
			player <- "backflip"
		}

		time.Sleep(30 * time.Millisecond)
	}
}

func player(name string, table chan Ball, coach chan string) {
	move := "forehand"

	for {
		select {
		case ball := <-table:
			fmt.Printf("Player %s %s'ing ball!\n", name, move)
			time.Sleep(100 * time.Millisecond)
			table <- ball
		case move = <-coach:
			fmt.Printf("Player %s is preparing to %s!\n", name, move)
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func main() {
	table := make(chan Ball)

	ping_coach, pong_coach := make(chan string), make(chan string)
	go coach(ping_coach)
	go coach(pong_coach)
	go player("Ping", table, ping_coach)
	go player("Pong", table, pong_coach)

	fmt.Println("Let the game begin!")
	table <- Ball{}
	time.Sleep(1 * time.Second)

	<-table
	close(table)
	fmt.Println("Game over!")
}
