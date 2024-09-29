package features

import "fmt"

// DemonstratePointers demonstrates the usage of pointers in Go.
// It includes examples of:
// - Declaring and initializing pointers
// - Printing pointer values and dereferencing pointers
// - Passing pointers to functions and modifying the original value
func DemonstratePointers() {
	// Using Pointers
	var intVar4 int = 4
	var intVar4Pointer *int = &intVar4

	fmt.Println(intVar4)
	fmt.Println(intVar4Pointer)

	// Dereferencing pointers
	fmt.Println(*intVar4Pointer)

	// Passing pointers to functions
	increment(intVar4Pointer)
	fmt.Println(intVar4)

	// Setting values of pointers
	var intVar5 *int = new(int)
	var intVar6 int

	fmt.Printf("The value of intVar5 is %v\n", *intVar5)
	fmt.Printf("The value of intVar6 is %v\n", intVar6)

	intVar5 = &intVar6
	*intVar5 = 6

	fmt.Printf("The value of intVar5 is %v\n", *intVar5)
	fmt.Printf("The value of intVar6 is %v\n", intVar6)

}

// increment takes a pointer to an integer and increments the value it points to by one.
func increment(number *int) {
	*number++
}
