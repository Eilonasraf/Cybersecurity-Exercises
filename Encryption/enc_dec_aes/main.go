package main

import (
	"bufio"
	"crypto/aes"    // AES encryption
	"crypto/cipher" // Cipher modes (like GCM)
	"crypto/rand"   // Random number generator for nonce
	"fmt"
	"io" // user_data/output for reading random bytes
	"os"
)

// Encrypt data using AES-GCM encryption
// Parameters:
// - data: the plaintext data to be encrypted (as a byte slice)
func encrypt(user_data, enc_key []byte) ([]byte, []byte, error) {
	// Step 1: Create an AES cipher block using the provided enc_key
	block, err := aes.NewCipher(enc_key)
	if err != nil {
		return nil, nil, err // If there is an erorr Return nils for results and err for the error
		                     // No ciphertext or nonce is returned if there's an error
	}

	// Step 2: Create a GCM cipher mode instance based on the AES block
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, nil, err // Return an error if GCM initialization fails
	}

	// Step 3: Generate a nonce (a unique number for this encryption)
	nonceSize := gcm.NonceSize()
	nonce := make([]byte, nonceSize) // Nonce size is determined by the GCM mode
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, nil, err // Return an error if GCM initialization fails
	}

	// Step 4: Encrypt the user_data using Seal.
	// Seal encrypts the plaintext and
	// appends the ciphertext to a destination slice (empty in this case).
	ciphertext := gcm.Seal(nil, nonce, user_data, nil)

	// Step 5: Return the encrypted user_data (ciphertext) and the nonce (needed for decryption)
	return ciphertext, nonce, nil
}

// Decrypt user_data that was encrypted using AES-GCM
// Parameters:
// - ciphertext: the encrypted user_data (as a byte slice)
// - enc_key: the same encryption enc_key used during encryption
// - nonce: the unique nonce generated during encryption (required for decryption)

func decrypt(ciphertext, enc_key, nonce []byte) ([]byte, error) {
	// Step 1: Create an AES cipher block using the provided enc_key
	block, err := aes.NewCipher(enc_key)
	if err != nil {
		return nil, err // Return an error if the enc_key size is invalid
	}

	// Step 2: Create a GCM cipher mode instance based on the AES block
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err // Return an error if GCM initialization fails
	}

	// Step 3: Decrypt the user_data using Open.
	// Open decrypts the ciphertext and appends the decrypted user_data (plaintext)
	// to a destination slice.
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err // Return an error if decryption fails
		// (e.g., wrong enc_key or tampered user_data)
	}

	// Step 4: Return the decrypted plaintext
	return plaintext, nil
}

func main() {

	// Define the user_data to be encrypted (plaintext). In Go, a string can be
	// converted to a byte slice using []byte().

	// Ask user for input to encrypt
	// Create a new reader

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the message you want to encrypt:")
	input, _ := reader.ReadString('\n')
	user_data := []byte(input)

	// key len options
	const (
		AES128 = 16
		AES192 = 24
		AES256 = 32
	)

	// Let user choose key len
	var encKey_choise int
	fmt.Println(
		"Select encryption key length (1 for 16 bytes (AES-128), " +
			"2 for 24 bytes (AES-192), " +
			"3 for 32 bytes (AES-256)):")
	fmt.Scan(&encKey_choise)

	// Define key based on choise
	var enc_key []byte

	switch encKey_choise {
	case 1:
		enc_key = []byte("thisis16byteskey") // 16 bytes for AES-128 (128-bit key)
	case 2:
		enc_key = []byte("thisis24byteslongpasskey") // 24 bytes for AES-192 (192-bit key)
	case 3:
		enc_key = []byte("thisis32byteslongpasskeyphrase") // 32 bytes for AES-256 (256-bit key)
	default:
		fmt.Println("Invalid choice, using 16 bytes AES-128 by default.")
		enc_key = []byte("thisis16byteskey")
	}

	// Step 1: Encrypt the user_data using the provided enc_key
	ciphertext, nonce, err := encrypt(user_data, enc_key)
	if err != nil {
		fmt.Println("Error during encryption:", err)
		return
	}

	// Print the ciphertext (the encrypted user_data) and the nonce
	fmt.Printf("Ciphertext (hex): %x\n", ciphertext)
	fmt.Printf("Nonce (hex): %x\n", nonce)

	// Step 2: Decrypt the ciphertext using the same enc_key and nonce
	decrypted_user_data, err := decrypt(ciphertext, enc_key, nonce)
	if err != nil {
		fmt.Println("Error during decryption:", err)
		return
	}

	// Print the decrypted message (should be the same as the original plaintext)
	fmt.Printf("Decrypted message: %s\n", decrypted_user_data)
}
