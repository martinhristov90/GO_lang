package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Declaring own type of map
	db := database{
		"shoes": 50,
		"socks": 5,
	}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list)) // Handles all calls to /list
	mux.Handle("/price", http.HandlerFunc(db.price)) // Handles all calls /price

	http.ListenAndServe("localhost:8000", mux)
}

// Own type for the prices, it has only String() method to override the %s
type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f USD", d)
}

// Main DB that contains KV for items and their prices
type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s : %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%s", "No such item")
		return
	}
	fmt.Fprintf(w, "The price of the requested item is %s", price)
}
