package main

import (
	"fmt"
	"net/http"
	"time"
	"log"
	"os"
)

func main() {

	err := WaitForServer(os.Args[1:])
	if err != nil {
		// function is going to try to reach the site exponationally, this logic is handle by the function itself,
		// if it returns nil, obviously the site cannot be reached
		fmt.Fprintf(os.Stderr, "Site is obviously down %v\n", err)
	}

}

func WaitForServer(url []string) error {
	const timeout = 1 * time.Minute

	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url[0])
		if err == nil {
			return nil
		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries))
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
