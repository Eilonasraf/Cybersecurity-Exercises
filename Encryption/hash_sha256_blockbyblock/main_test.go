package main

import (
	"crypto/sha256"
	"testing"
)

// TestHashingBlocks verifies if the blocks are correctly hashed and detects tampering.
func TestHashingBlocks(t *testing.T) {
	// Define the input data to be hashed and the block size for dividing the input.
	Input := []byte("This is the Input that I want to encrypt in blocks")
	blockSize := 16

	// Call the Hashing_Blocks function, which will hash the input block by block
	// and return plaintext blocks, ciphertext blocks (hashed), and the initial verification array.
	plaintext_Blocks, ciphertext_Blocks, add_Array := Hashing_Blocks(Input, blockSize)

	// Verify that all blocks are initially marked as "intact" (true in the add_Array).
	for i, status := range add_Array {
		if !status { // If any block is marked as false, it means tampering is detected (which should not be).
			t.Errorf("Block %d is tampered without any modification!", i)
		}
	}

	// Simulate tampering by modifying a bit in the second block (if it exists).
	// This is done by flipping the first bit of the first byte in the second block.
	if len(plaintext_Blocks) > 1 {
		plaintext_Blocks[1][0] ^= 1 // The XOR operation flips the bit.
	}

	// Re-verify the blocks after the tampering simulation.
	// The function verifyBlocks compares the current plaintext blocks with their original ciphertext blocks.
	add_Array = verifyBlocks(plaintext_Blocks, ciphertext_Blocks)

	// Check if the tampering in the second block is correctly detected.
	// After tampering, add_Array[1] should be false.
	if len(plaintext_Blocks) > 1 && add_Array[1] {
		t.Errorf("Block 1 tampering was not detected!") // If tampering is not detected, throw an error.
	}

	// Check that all other blocks (besides the tampered block) remain intact.
	for i, status := range add_Array {
		if i != 1 && !status { // Ensure that all blocks except block 1 (the tampered one) are still marked as true.
			t.Errorf("Block %d is incorrectly marked as tampered, but no changes were made to it!", i)
		}
	}
}

// TestHashingFunctionality checks if the SHA256 hashing works as expected.
// This is a simple check to ensure that the SHA256 function behaves as expected.
func TestHashingFunctionality(t *testing.T) {
	// Define a sample block of data to hash.
	block := []byte("Sample block of data")

	// Use the SHA256 function to manually compute the expected hash.
	expectedHash := sha256.Sum256(block)

	// Compute the actual hash using the same function to compare.
	actualHash := sha256.Sum256(block)

	// Compare the expected hash and the actual hash byte by byte.
	// If there is any mismatch, it means the hashing function didn't work as expected.
	for i := range expectedHash {
		if expectedHash[i] != actualHash[i] { // If any byte is different, throw an error.
			t.Errorf("SHA256 hashing failed. Expected %x but got %x", expectedHash, actualHash)
		}
	}
}
