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

}
