

package main

import (
	"fmt"
	"sync"
)

func golangfans(ch8 chan interface{}) {

	fmt.Println(" everyday is now monday")

	var wg sync.WaitGroup
	wg.Add(1)

	fmt.Println(" waiting queue is now added ")

	func(ch8 chan interface{}) {

		defer wg.Done()
		fmt.Println(" sunday is now monday")
		select {
		case a := <-ch8:
			fmt.Println(" sunday is now monday", a)
		}

	}(ch8)

	fmt.Println(" golangfans end")

}

func sum(ch8 chan interface{}) {

	fmt.Println(" sum goroutine starts")

	ch8 <- "happiness"

	fmt.Println(" sum goroutine ends ")
}

func main() {

	ch8 := make(chan interface{})

	go sum(ch8) // register gororutine

	golangfans(ch8)

}





case 2:



package main

import (
	"fmt"
)


func golangfans(){

     fmt.Println(" everyday is now monday")

}

func main() {


	ch8 := make(chan interface{})

	func() {

		fmt.Println(" sunday is now monday")

		select {
		case a := <-ch8:
			fmt.Println(" sunday is now monday", a)
		}

	}()


           golangfans()

}