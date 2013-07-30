package main

import (
	"fmt"
	"time"
)

type Ball struct{}

func main() {
	table := make(chan Ball)
	go player("ping", table)
	go player("pong", table)

	fmt.Println("Let the game begin!")
	table <- Ball{}
	time.Sleep(1 * time.Second)

	<-table
	close(table)
	fmt.Println("Game over!")

	time.Sleep(100 * time.Millisecond)

	panic("print stack")
}

func player(name string, table chan Ball) {
	for ball := range table {
		fmt.Println(name)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}
