package features

import "fmt"

// DemonstrateMaps demonstrates various operations on a map in Go.
// It includes the following operations:
// - Initializing a map with string keys and int32 values.
// - Printing the entire map.
// - Accessing a map value by its key and checking if the key exists.
// - Deleting a key-value pair from the map.
// - Iterating over the map to print all keys.
// - Iterating over the map to print all key-value pairs.
func DemonstrateMaps() {
	// Maps
	var intMap map[string]int32 = map[string]int32{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
	}

	fmt.Println(intMap)

	demonstrateMapIndex(intMap)

	var num, found = intMap["six"] // Check if key exists in map, will return 2 values, value and a boolean

	// Check if key exists in map
	if found {
		fmt.Println(num)
	} else {
		fmt.Println("Not found")
	}

	delete(intMap, "five") // Built-in function to delete key-value pair from map

	fmt.Println(intMap)

	// For loop returning the keys of the map
	for name := range intMap {
		fmt.Println(name)
	}

	// For loop returning the keys and values of the map
	for name, num := range intMap {
		fmt.Println(name, num)
	}
}

// demonstrateMapIndex prints the values associated with specific keys in a given map.
// It expects a map with string keys and int32 values as its parameter.
// The function prints the values for the keys "one", "two", "three", "four", and "five".
func demonstrateMapIndex(intMap map[string]int32) {
	fmt.Println(intMap["one"])
	fmt.Println(intMap["two"])
	fmt.Println(intMap["three"])
	fmt.Println(intMap["four"])
	fmt.Println(intMap["five"])
}
