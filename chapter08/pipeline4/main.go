package main

import (
	"fmt"
)

func main() {
	numbers := make(chan int)
	squaredNumbers := make(chan int)
	go generator(numbers)
	go squarer(numbers, squaredNumbers)
	printer(squaredNumbers)

}

func generator(c chan<- int) {
	for i := 1; i <= 100; i++ {
		c <- i
	}
	close(c)
}

func squarer(recv <-chan int, send chan<- int) {
	for i := range recv {
		send <- i * i
	}
	close(send)
}

func printer(recv <-chan int) {
	for i := range recv {
		fmt.Println(i)
	}
}
