package features

import "fmt"

// DemonstrateFunctions demonstrates the declaration, definition, and calling of functions in Go.
// It covers:
// - Function declaration and definition
// - Function call with multiple return values
// - Error handling in functions
func DemonstrateFunctions() (int, int, error) {
	// Functions
	printMe("Hello, World! 4")

	var numerator int = 11
	var denominator int = 2
	var result, remainder, err = divide(numerator, denominator) // Function call, multiple return values, error handling

	return result, remainder, err
}

// printMe prints the provided string value to the standard output.
//
// Parameters:
//   - printValue: The string value to be printed.
func printMe(printValue string) {
	fmt.Println(printValue)
}

// divide performs integer division of the numerator by the denominator.
// It returns the quotient, the remainder, and an error if the denominator is zero.
//
// Parameters:
//
//	numerator - the integer value to be divided.
//	denominator - the integer value by which the numerator is divided.
//
// Returns:
//
//	quotient - the result of the division.
//	remainder - the remainder of the division.
//	err - an error if the denominator is zero, otherwise nil.
func divide(numerator int, denominator int) (int, int, error) {
	var err error

	if denominator == 0 {
		err = fmt.Errorf("division by zero")
		return 0, 0, err
	}

	quotient := numerator / denominator
	remainder := numerator % denominator
	return quotient, remainder, err
}
