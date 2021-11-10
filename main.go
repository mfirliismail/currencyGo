package main

import (
	"fmt"
	"sync"
	"time"
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

	// wgWait()
	// channelStudy()
	// bufferedChannel()

	ping := make(chan int)
	pong := make(chan int)

	go pinger(ping, pong)
	go ponger(ping, pong)

	pong <- 1

	select {}

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

type Car struct {
	Model string
	Age   int
}

func bufferedChannel() {
	chanel := make(chan *Car, 3)

	go func() {
		chanel <- &Car{"BMW", 2001}
		chanel <- &Car{"Daihatsu", 1998}
		chanel <- &Car{"Honda JH221", 2014}
		chanel <- &Car{"Pajero 98K1", 2021}

		close(chanel)
	}()

	for i := range chanel {
		fmt.Println(i.Model, i.Age)
	}
}
func pinger(pinger <-chan int, ponger chan<- int) {
	for {
		<-pinger
		fmt.Println("ping")
		time.Sleep(time.Second)
		ponger <- 1
	}
}

func ponger(pinger chan<- int, ponger <-chan int) {
	for {
		<-ponger
		fmt.Println("pong")
		time.Sleep(time.Second)
		pinger <- 1
	}
}
