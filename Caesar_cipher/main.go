package main

import (
	"fmt"
	"unicode"
)

// Caesar cipher function
// This function takes a string `text` and an integer `shift` as inputs and returns
// an encrypted or decrypted string depending on the sign of the `shift`.
func caesarCipher(text string, shift int) string {
	// Create a rune slice with the same length as the input text.
	// Rune is used because it can handle characters that may have multi-byte representations.
	res := make([]rune, len(text))

	// Iterate through each character (rune) in the input string.
	for i, char := range text {
		// Check if the character is a letter (either uppercase or lowercase)
		if unicode.IsLetter(char) {
			// Determine the base character (either 'A' or 'a') depending on whether
			// the current character is uppercase or lowercase.
			base := 'A'
			if unicode.IsLower(char) {
				base = 'a'
			}

			// Apply the Caesar cipher formula:
			// 1. Subtract the base value from the current character to normalize it to a 0-based index.
			// 2. Add the shift value.
			// 3. Use modulo 26 to wrap around the alphabet if needed.
			// 4. Add the base value back to convert it to the correct letter.
			Caesar_formula := (int(char-base)+shift)%26 + int(base)

			// Assign the calculated character to the result slice.
			res[i] = rune(Caesar_formula)
		} else {
			// If the character is not a letter, keep it unchanged.
			// For example, spaces, punctuation marks, and numbers are unaffected.
			res[i] = char
		}
	}
	// Convert the rune slice back into a string and return it.
	return string(res)
}

func main() {
	// Define the plaintext message to be encrypted.
	plaintext := "Hello, World!"

	// Define the number of positions to shift for the Caesar cipher.
	// A positive shift encrypts, while a negative shift decrypts.
	shift := 3

	// Encrypt the plaintext using the Caesar cipher.
	encrypted := caesarCipher(plaintext, shift)

	// Decrypt the encrypted message by applying the negative of the original shift.
	decrypted := caesarCipher(encrypted, -shift)

	// Print the original plaintext, encrypted message, and decrypted message.
	fmt.Printf("Plaintext: %s\n", plaintext)
	fmt.Printf("Encrypted: %s\n", encrypted)
	fmt.Printf("Decrypted: %s\n", decrypted)
}