
Case 1: 


package main

import (
	"fmt"
)

func printIntegers(done <-chan struct{}, intStream <-chan int) {

	fmt.Println("printIntegers goroutine ")

	for {

		fmt.Println("infinite  loop ")
		select {
		case i := <-intStream:
			fmt.Println("int data existed")
			fmt.Println(i)
		case <-done:
			return
		}
	}

	fmt.Println("end of infinite loop ")

}

func main() {

	var chance chan struct{} = make(chan struct{})
	var chance1 chan int = make(chan int)
	go printIntegers(chance, chance1)
	chance1 <- 1
	fmt.Println("main existed")

}


Case 2:

package main

import (
	"fmt"
)

func printIntegers(done <-chan struct{}, intStream <-chan int) {

	fmt.Println("printIntegers goroutine ")

	for {

		fmt.Println("infinite  loop ")
		select {
		case i := <-intStream:
			fmt.Println("int data existed")
			fmt.Println(i)
		case <-done:
			fmt.Println("struct data existed")
			return
		}
	}

	fmt.Println("end of infinite loop ")

}

func main() {

	var chance chan struct{} = make(chan struct{})
	var chance1 chan int = make(chan int)
	go printIntegers(chance, chance1)
	type coder struct {
	}
	chance <- coder{}
	fmt.Println("main existed")

}

Case 3: 

package main

import (
	"fmt"
	"context"
)


func printIntegers( ctx context.Context , intStream <-chan int) {

  for{
    select {
      case i := <-intStream:
        fmt.Println(i)
      case <-ctx.Done():
          fmt.Println("context done")
        return
    }
  }

}

func main() {

         ctx := context.TODO()
	var chance1 chan int = make(chan int)
	go printIntegers(ctx, chance1)

	chance1 <- 1
	fmt.Println("main existed")

}