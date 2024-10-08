package main

import (
	"crypto/sha256"
	"fmt"
)

// Encrypts input block by block using SHA256, returns plaintext blocks, ciphertext blocks, and a verification array.
func Hashing_Blocks(Input []byte, blockSize int) ([][]byte, [][]byte, []bool) {
	var plaintext_Blocks [][]byte  // Stores the original data blocks.
	var ciphertext_Blocks [][]byte // Stores the hashed (ciphertext) blocks.
	var verificationTrack []bool   // Tracks whether each block is initially verified (all set to true).

	// Loop through the input data, processing one block at a time.
	for i := 0; i < len(Input); i += blockSize {
		The_end := i + blockSize
		// If we reach the last block and it's smaller than blockSize, adjust the end index.
		if The_end > len(Input) {
			The_end = len(Input)
		}

		// Extract the current block from the input data.
		curr_block := Input[i:The_end]

		// Hash the current block using SHA256.
		curr_hash := sha256.Sum256(curr_block)

		// Append the block and its corresponding hash to the respective arrays.
		plaintext_Blocks = append(plaintext_Blocks, curr_block)
		ciphertext_Blocks = append(ciphertext_Blocks, curr_hash[:])

		// Initially, set verificationTrack to true for all blocks, assuming they are intact.
		verificationTrack = append(verificationTrack, true)
	}

	return plaintext_Blocks, ciphertext_Blocks, verificationTrack
}

// Verifies if the blocks have been tampered with by comparing the plaintext to the hashed values in ciphertext.
func verifyBlocks(plaintext_Blocks, ciphertext_Blocks [][]byte) []bool {
	verificationTrack := make([]bool, len(plaintext_Blocks)) // Initialize an array to track block verification.

	// Loop through each plaintext block and verify if its hash matches the original ciphertext.
	for i, block := range plaintext_Blocks {
		curr_hash := sha256.Sum256(block) // Recalculate the hash for the current block.
		is_equal := true

		// Compare the current hash with the original ciphertext (hashed block).
		if len(curr_hash) != len(ciphertext_Blocks[i]) {
			is_equal = false
		}
		for j := range curr_hash {
			if curr_hash[j] != ciphertext_Blocks[i][j] {
				is_equal = false
				break
			}
		}
		// Update verification status: true if intact, false if tampered.
		verificationTrack[i] = is_equal
	}

	return verificationTrack
}

func main() {
	// Input data to be processed (converted to a byte slice).
	Input := []byte("This is the Input that i want to encrypt in blocks")

	// Define the block size for splitting the input data.
	blockSize := 16

	// Encrypt the data block by block and get the initial verification array.
	plaintext_Blocks, ciphertext_Blocks, add_Array := Hashing_Blocks(Input, blockSize)

	// Print the original data, plaintext blocks, ciphertext blocks (hashes), and the initial verification array.
	fmt.Println("Original data:", string(Input))
	fmt.Println("Plaintext Blocks:", plaintext_Blocks)
	fmt.Println("Ciphertext Blocks (SHA256 Hashes):", ciphertext_Blocks)
	fmt.Println("Initial Verification Array:", add_Array)

	// Verify the integrity of the blocks to check if they’ve been altered.
	add_Array = verifyBlocks(plaintext_Blocks, ciphertext_Blocks)

	// Print the updated verification array after checking for tampered blocks.
	fmt.Println("Updated Verification Array:", add_Array)

	// Check for tampered blocks and print a message for each block's status.
	for i, status := range add_Array {
		if !status {
			fmt.Printf("Block %d has been tampered with!\n", i)
		} else {
			fmt.Printf("Block %d is intact.\n", i)
		}
	}
}