package features

import "fmt"

// DemonstrateVariables demonstrates the declaration, assignment, and printing of various types of variables in Go.
// It covers:
// - Declaring variables without initialization
// - Assigning values to variables
// - Printing variables
// - Declaring and assigning variables in a single statement
// - Using type inference to declare and assign variables
//
// The function handles the following types:
// - int
// - float64
// - string
// - bool
// - rune
func DemonstrateVariables() {
	// Declaring variables
	var intVar int
	var floatVar float64
	var stringVar string
	var boolVar bool
	var char rune

	// Assigning values to variables
	intVar = 1
	floatVar = 1.1
	stringVar = "Hello, World!"
	boolVar = true
	char = 'a'

	// Printing variables
	fmt.Println(intVar)
	fmt.Println(floatVar)
	fmt.Println(stringVar)
	fmt.Println(boolVar)
	fmt.Println(char)
	fmt.Println(string(char))

	// Declaring and assigning variables
	var intVar2 int = 2
	var floatVar2 float64 = 2.2
	var stringVar2 string = "Hello, World! 2"
	var boolVar2 bool = false
	var char2 rune = 'b'

	// Printing variables
	fmt.Println(intVar2)
	fmt.Println(floatVar2)
	fmt.Println(stringVar2)
	fmt.Println(boolVar2)
	fmt.Println(char2)

	// Declaring and assigning variables with type inference
	intVar3 := 3
	floatVar3 := 3.3
	stringVar3 := "Hello, World! 3"
	boolVar3 := true
	char3 := 'c'

	// Printing variables
	fmt.Println(intVar3)
	fmt.Println(floatVar3)
	fmt.Println(stringVar3)
	fmt.Println(boolVar3)
	fmt.Println(char3)
}
