package features

import "fmt"

// DemonstrateArrays demonstrates various operations on arrays and slices in Go.
// It includes examples of array initialization, indexing, slicing, and memory location printing.
// The function performs the following tasks:
// 1. Initializes an array of integers and prints it.
// 2. Demonstrates different ways to index and slice the array, printing the results.
// 3. Prints the memory locations of the array and its elements.
// 4. Initializes another array using a different syntax and prints it.
// 5. Creates a slice from the array and prints it.
// 6. Appends an element to the slice and prints the updated slice.
func DemonstrateArrays() {
	// Arrays
	var intArray [5]int32
	intArray[0] = 1
	intArray[1] = 2
	intArray[2] = 3
	intArray[3] = 4
	intArray[4] = 5

	fmt.Println(intArray)

	// Indexing arrays
	fmt.Println(intArray[0])               // First element
	fmt.Println(intArray[1:5])             // Second to fifth element - means start from second element and go upto fifth element
	fmt.Println(intArray[:5])              // First to fifth element - means start from first element and go upto fifth element
	fmt.Println(intArray[1:])              // Second to last element - means start from second element and go upto last element
	fmt.Println(intArray[:])               // All elements
	fmt.Println(intArray[len(intArray)-1]) // Last element

	// Memory location of array and its elements
	fmt.Println(&intArray)
	fmt.Println(&intArray[0])
	fmt.Println(&intArray[1])
	fmt.Println(&intArray[2])
	fmt.Println(&intArray[3])
	fmt.Println(&intArray[4])

	// Memory allocations
	var intArray2 = [...]int32{1, 2, 3, 4, 5}
	fmt.Println(intArray2)

	// Slices
	var intSlice []int32 = intArray[1:4]
	fmt.Println(intSlice)

	// Append to slice
	intSlice = append(intSlice, 6)
	fmt.Println(intSlice)
}
