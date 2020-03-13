package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Declaring own type of map
	db := database{
		"shoes": 50,
		"socks": 5,
	}
	// Actual serving of the page
	log.Fatal(http.ListenAndServe("localhost:8000", db))

}

// Own type for the prices, it has only String() method to override the %s
type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

// Main DB that contains KV for items and their prices
type database map[string]dollars

// Method called by http.ListenAndServe
func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path { // switching on this Path string from URL struct, it is relative path
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s : %s\n", item, price) // for price, it calls String() of dollars type
		}
	case "/price":
		item := req.URL.Query().Get("item") // Getting the item parameter from GET request
		price, ok := db[item]               // Check if we have it in our map of items
		if !ok {                            //return error
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "No such item %q\n", item) // uses Write() of http.ResponseWrite interface
			return
		}
		fmt.Fprintf(w, "%s\n", price) // return the price of the item

	default:
		w.WriteHeader(http.StatusNotFound) // return error if none of case matches
		fmt.Fprintf(w, "no such page %q", req.URL)
	}
}
