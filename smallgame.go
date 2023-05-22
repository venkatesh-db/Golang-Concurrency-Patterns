package main

import (
	"fmt"
	"time"
)

type Ball struct{ hits int }

func player(name string, table chan *Ball) {

	fmt.Println("player")
	for {
		ball := <-table // player grabs the ball
		ball.hits++
		fmt.Println(name, ball.hits)
		time.Sleep(3 * time.Millisecond)
		table <- ball // pass the ball
	}
}

func main() {

	table := make(chan *Ball)

	go player("ping", table)
	go player("pong", table)

	fmt.Println("writing data in to channel ")

	table <- new(Ball) // game on; toss the ball

	fmt.Println("main sleep")

	time.Sleep(1 * time.Second)

	<-table // game over, grab the ball

	//panic("show me the stack")

	fmt.Println(" main over")

}
