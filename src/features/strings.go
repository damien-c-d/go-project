package features

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// DemonstrateStrings demonstrates various string operations in Go.
// It performs the following tasks:
// 1. Initializes a string variable and prints its value.
// 2. Accesses and prints a specific byte from the string.
// 3. Iterates over the string and prints each character with its index.
// 4. Prints the length of the string in bytes and characters.
// 5. Builds a new string using a custom function and prints it.
// 6. Builds another new string using strings.Builder and prints it.
func DemonstrateStrings() {
	var value = "Résumé"
	fmt.Println(value)

	var indexed = value[1]
	fmt.Printf("%v %T", indexed, indexed)
	for i, v := range value {
		fmt.Println(i, v)
	}
	fmt.Printf("\nThe length of the string is %d", len(value))    // Prints length in bytes
	fmt.Printf("\nThe length of the string is %d", strLen(value)) // Prints length in characters

	var newValue = buildString()
	fmt.Printf("\nThe new value is %v", newValue)

	var newValue2 = buildStringWithStringsBuilder()
	fmt.Printf("\nThe new value is %v\n", newValue2)
}

// strLen returns the number of runes in the given string.
// It accurately counts the length of the string by considering
// multi-byte characters, ensuring the correct length is returned.
//
// Parameters:
//   - str: The input string whose length is to be calculated.
//
// Returns:
//   - int: The number of runes in the input string.
func strLen(str string) int {
	// Get the actual length of the string
	return utf8.RuneCountInString(str)
}

// buildString concatenates a slice of strings into a single string.
// It iterates over a predefined slice of strings and appends each element
// to a result string, which is then returned.
func buildString() string {
	var strSlice = []string{"r", "é", "s", "u", "m", "é"}

	var catStr = ""
	for i := range strSlice {
		catStr += strSlice[i]
	}

	return catStr
}

// buildStringWithStringsBuilder constructs a string by concatenating elements
// from a slice of strings using the strings.Builder for efficient string
// concatenation. It returns the resulting concatenated string.
func buildStringWithStringsBuilder() string {
	var strSlice = []string{"r", "é", "s", "u", "m", "é"}

	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}

	return strBuilder.String()
}
