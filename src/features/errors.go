package features

import "fmt"

// DemonstrateErrorHandling prints the result of a division operation based on the provided values.
// If an error is present, it prints the error message. Otherwise, it prints the quotient and remainder.
func DemonstrateErrorHandling(err error, remainder int, result int) {
	// Error handling
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else if remainder == 0 {
		fmt.Printf("Quotient: %d\n", result)
	} else {
		fmt.Printf("Quotient: %d, Remainder: %d\n", result, remainder)
	}
}
