package main

import "fmt"

func main() {

	ages := make(map[string]int)

	ages["Marti"] = 29
	ages["Kremi"] = 25


	name , exists := ages["Marti"]
	// If the key does not exists in the map, default type of the Value is return, in this case 0 (int), belowe is the statement check if the key exists in the map.
	fmt.Printf("%v %v\n",name,exists); if !exists { fmt.Printf ("The key does not exists\n")}

}
