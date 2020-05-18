package main

import (
	"fmt"
	"net/http"
	"log"
	"sync"
)

var urls = []string{
	"https://google.com",
	"https://twitter.com",
	"https://tugab.bg",
}

func fetchStatus(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Fprintf(w, "%+v\n", err)
			}
			fmt.Fprintf(w, "%+v\n", resp.StatusCode)
			wg.Done()
		}(url)
	}
	wg.Wait()
}

func main() {
	http.HandleFunc("/", fetchStatus)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
