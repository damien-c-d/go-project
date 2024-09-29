package features

import "fmt"

// DemonstrateSlices showcases the usage of slices in Go. It includes:
// - Creating and copying slices
// - Modifying slices and observing changes in the original slice
// - Passing arrays to functions and observing the behavior
// - Passing arrays to functions using pointers and observing the behavior
//
// The function demonstrates how slices share the same underlying array,
// leading to changes in one slice reflecting in another. It also shows
// how arrays are passed by value to functions, and how using pointers
// can allow functions to modify the original array.
func DemonstrateSlices() {

	// Demonstrating slices
	var slice = []int32{1, 2, 3, 4, 5}
	var sliceCopy = slice // Copying slice to another variable
	sliceCopy[2] = 4
	fmt.Println(slice)
	fmt.Println(sliceCopy)

	// Passing slices to functions
	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("The memory address of thing1 is %p\n", &thing1)
	var result = square(thing1)
	fmt.Printf("Thing1 is %v\n", thing1)
	fmt.Printf("The result is %v\n", result)

	// Passing slices to functions using pointers
	var thing3 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("The memory address of thing2 is %p\n", &thing3)
	var result2 = squarePtr(&thing3)
	fmt.Printf("Thing3 is %v\n", thing3)
	fmt.Printf("The result is %v\n", result2)

}

// square takes an array of 5 float64 elements, squares each element,
// and returns the resulting array. It also prints the memory address
// of the input array.
//
// Parameters:
// - thing2: An array of 5 float64 elements.
//
// Returns:
// - An array of 5 float64 elements where each element is the square
//   of the corresponding element in the input array.
func square(thing2 [5]float64) [5]float64 {
	fmt.Printf("The memory address of thing2 is %p\n", &thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}

	return thing2
}

// squarePtr takes a pointer to an array of 5 float64 values, squares each element in place,
// and returns the modified array. It also prints the memory address of the input array.
//
// Parameters:
//   - thing3: A pointer to an array of 5 float64 values.
//
// Returns:
//   - A copy of the modified array with each element squared.
func squarePtr(thing3 *[5]float64) [5]float64 {
	fmt.Printf("The memory address of thing3 is %p\n", thing3)
	for i := range thing3 {
		thing3[i] = thing3[i] * thing3[i]
	}

	return *thing3
}
