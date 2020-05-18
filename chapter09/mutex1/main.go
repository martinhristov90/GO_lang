package main

import (
	"fmt"
	"sync"
)

var (
	balance int
	mutex   sync.Mutex
)

func accBalance(wg *sync.WaitGroup) {
	//defer wg.Done()

	mutex.Lock()
	b := balance
	fmt.Println("Current balance", b)
	mutex.Unlock()
}

func deposit(dep int, wg *sync.WaitGroup) {

	defer wg.Done()

	mutex.Lock()
	balance += dep
	mutex.Unlock()

}

func main() {

	var wg sync.WaitGroup

	wg.Add(2)
	balance = 100

	go deposit(200, &wg)
	go deposit(200, &wg)

	accBalance(&wg) // balance 100

	wg.Wait()

	accBalance(&wg) // balance 500


}
