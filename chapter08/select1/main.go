package main

import (
	"fmt"
	"math/rand"
	"time"
)

var ch1 = make(chan string)
var ch2 = make(chan string)

func main() {

	go worker(ch1, "channel1")
	go worker(ch2, "channel2")

	for {
		select {
		case work := <-ch1:
			fmt.Println("The worker did some work and send it to", work)
		case work := <-ch2:
			fmt.Println("The worker did some work and send it to", work)
		}
	}

}

func worker(ch chan<- string, chName string) {
	for {
		rand.Seed(time.Now().UnixNano())
		rndTime := rand.Intn(10)
		//fmt.Printf("sleeping for %d seconds\n", rndTime)
		time.Sleep(time.Duration(rndTime) * time.Second)
		ch <- fmt.Sprintf("%s, the work took %d seconds", chName, rndTime) // sending the work done :)
	}
}
