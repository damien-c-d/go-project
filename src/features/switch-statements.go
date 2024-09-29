package features

import "fmt"

// DemonstrateSwitchStatements demonstrates the usage of switch statements
// by calling two different functions: switchStatement and conditionalSwitchStatement.
// It takes two integer parameters:
// - result: an integer value to be used in the switch statements.
// - remainder: an integer value to be used in the switch statements.
func DemonstrateSwitchStatements(result int, remainder int) {
	switchStatement(result, remainder)
	conditionalSwitchStatement(result, remainder)
}

// switchStatement prints the quotient and remainder based on the provided result and remainder.
// If the remainder is zero, it prints only the quotient. Otherwise, it prints both the quotient and the remainder.
//
// Parameters:
//   - result: The quotient of a division operation.
//   - remainder: The remainder of a division operation.
func switchStatement(result, remainder int) {
	switch {
	case remainder == 0:
		fmt.Printf("Quotient: %d\n", result)
	default:
		fmt.Printf("Quotient: %d, Remainder: %d\n", result, remainder)
	}
}

// conditionalSwitchStatement prints the quotient and optionally the remainder
// based on the value of the remainder. If the remainder is 0, it prints only
// the quotient. Otherwise, it prints both the quotient and the remainder.
//
// Parameters:
//   - result: The quotient of a division operation.
//   - remainder: The remainder of a division operation.
func conditionalSwitchStatement(result, remainder int) {
	switch remainder {
	case 0:
		fmt.Printf("Quotient: %d\n", result)
	default:
		fmt.Printf("Quotient: %d, Remainder: %d\n", result, remainder)
	}
}
