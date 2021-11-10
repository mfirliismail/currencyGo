package main

import (
	"fmt"
	"sync"
	// "time"
)

// func heavy() {
// 	for {

// 		time.Sleep(time.Second * 1)
// 		fmt.Println("Helo")
// 	}

// }
func wgWait() {
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func() {
		fmt.Println("Hello")
		wg.Done()
	}()
	go func() {
		fmt.Println("World")
		wg.Done()
	}()

	fmt.Println("First")
	wg.Wait()
	fmt.Println("Last")
}

func main() {

	wgWait()
	channelStudy()

}

func channelStudy() {
	c := make(chan int)
	// <name > chan <datatype>

	// send
	go func() {
		c <- 1
	}()

	//sniff
	val := <-c

	fmt.Println(val)
	// send
	go func() {
		c <- 2
	}()
	//sniff
	val = <-c
	fmt.Println(val)
	fmt.Println(c)
}
