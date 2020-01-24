package main

import "fmt"

func main() {

	ages1 := map[string]int{"Marti": 29, "Katya": 65, "Kremi": 25}
	ages2 := map[string]int{"Marti": 29, "Katya": 65, "Kremi": 23}

	fmt.Printf("%t\n", map_equal(ages1, ages2))
}

// map1, map2 map[string]int means that map1 map2 are the same type
func map_equal(map1, map2 map[string]int) bool {
	// No possibility both maps to be equal if their length is different, just return false
	if len(map1) != len(map2) {
		return false
	}
	// Iterating over map1 keys and values

	for map1_k, map1_v := range map1 {
		// Trying to see if map1_k(map1 key) is available in map2 keys, if it is not `ok` would be false (reverse it to be true)
		// map2_v contains the value of the map1_k if the key is present in map2,  `|| map2_v != map1_v` checks whether the values of matching keys in both maps are the same
		if map2_v, ok := map2[map1_k]; !ok || map2_v != map1_v {
			// if the map1_k is not present in map2 or the values are not the same return false
			return false
		}
	}
	// If both checks pass, guess they should be the same ...
	return true
}
