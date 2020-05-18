package main

// Guide here : https://krakensystems.co/blog/2019/golang-race-detection

import(
	"fmt"
	"sync"

)

var Cnt int
//var mu sync.Mutex

func Run(amount int, wg *sync.WaitGroup) {
	for i := 0; i < amount; i++ {
	  //mu.Lock() Fixes the race conditions
	  Cnt++
	  //mu.Unlock()
	}
	defer wg.Done()
  }

func main(){

	var wg sync.WaitGroup
	// wg:= &sync.WaitGroup{}

	for i := 0 ; i < 10 ; i++ {

		wg.Add(1)
		go func(){
			//defer wg.Done()
			Run(1000,&wg)
		}()
	}
	wg.Wait()


	fmt.Println(Cnt)
	fmt.Println(&Cnt)
	// No mutex result is 7585
	// With mutex it works as expected
	// go run -race main.go


	//==================
	//==================
	//WARNING: DATA RACE
	//Read at 0x000001228360 by goroutine 9:
	//  main.Run()
	//      /Users/marti/HashiCorpWork/Tasks/GO_lang/chapter09/race1/main.go:17 +0x47
	//  main.main.func1()
	//      /Users/marti/HashiCorpWork/Tasks/GO_lang/chapter09/race1/main.go:33 +0x41
	//
	//Previous write at 0x000001228360 by goroutine 7:
	//  main.Run()
	//      /Users/marti/HashiCorpWork/Tasks/GO_lang/chapter09/race1/main.go:17 +0x63
	//  main.main.func1()
	//      /Users/marti/HashiCorpWork/Tasks/GO_lang/chapter09/race1/main.go:33 +0x41
	//
	//Goroutine 9 (running) created at:
	//  main.main()
	//      /Users/marti/HashiCorpWork/Tasks/GO_lang/chapter09/race1/main.go:31 +0xab
	//
	//Goroutine 7 (finished) created at:
	//  main.main()
	//      /Users/marti/HashiCorpWork/Tasks/GO_lang/chapter09/race1/main.go:31 +0xab
	//==================
	//2911
	//0x1228360
	//Found 2 data race(s)
	//exit status 66

	// this 0x1228360 is the address where all goroutine are trying to read and write to and from
}