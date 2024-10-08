package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s) // Read the input string

	var shiftedMessage string

	// Loop over each character in the string
	for i := 0; i < len(s); i++ {
		// Convert the character to an integer value
		currentDigit := int(s[i] - '0') // Convert character to integer

		// Shift the digit cyclically (9 becomes 0, etc.)
		shiftedDigit := (currentDigit + 1) % 10

		// Append the shifted digit as a string to the result
		shiftedMessage += string(shiftedDigit + '0') // Convert integer back to character
	}

	// Output the shifted message
	fmt.Println(shiftedMessage)
}