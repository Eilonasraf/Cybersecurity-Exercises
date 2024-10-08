package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

// Generic error-handling function
// This function takes an error and a message as arguments.
// If an error occurred, it prints the message and exits the program.
func Generic_err_handling(err error, message string) {
	if err != nil {
		fmt.Println(message, err)
		os.Exit(1)
	}
}

// Function to hash a file block by block, while performing an integrity check on the entire file
// This function reads the file `inputFile` block by block, hashes each block, and performs
// an ongoing hash of the entire file to check its integrity.
func hashFileBlockByBlock_With_IntegrityCheck(inputFile string, blockSize int) {
	// Open the input file
	inFile, err := os.Open(inputFile)
	Generic_err_handling(err, "Error opening input file:")
	defer func() {
		closeErr := inFile.Close()
		if closeErr != nil {
			fmt.Println("Error closing input file:", closeErr)
		}
	}()

	// Create a buffer (byte slice) of size blockSize
	// It acts as a temporary storage space for parts of the file reading.
	buffer := make([]byte, blockSize)
	var fullFileHash [32]byte // Store the hash for the entire file (final integrity hash)
	// 256 bits equals 32 bytes (since 1 byte = 8 bits, so 256 bits / 8 bits per byte = 32 bytes).

	// File to save block hashes
	hashesFile, err := os.Create("hashed_blocks.txt")
	Generic_err_handling(err, "Error creating hash file:")
	defer hashesFile.Close()

	// The file is read block by block into this buffer.
	for {
		// Read a block of data from the file into the buffer
		bytesRead, err := inFile.Read(buffer)
		// bytesRead tells how many bytes were actually read during this read operation

		if err != nil && err != io.EOF { // io.EOF means the end of the file was reached
			Generic_err_handling(err, "Error reading input file block:")
		}
		if bytesRead == 0 {
			// No more data available to read, we've reached the end of the file.
			break
		}

		// Get the current block, slice the buffer to contain only the valid data
		curr_Block := buffer[:bytesRead]

		// Hash the block using SHA256
		hashed_block := sha256.Sum256(curr_Block)

		// Print the hash of the current block in hex format
		fmt.Printf("Hashed block: %x\n", hashed_block)

		// Write the block hash to the output file (hashed_blocks.txt)
		hashesFile.WriteString(fmt.Sprintf("%x\n", hashed_block))

		// Combine the full file hash with the current block hash to create an ongoing hash
		combined_HashInput := append(fullFileHash[:], hashed_block[:]...)

		// Hash the new slice to update the full file integrity hash
		fullFileHash = sha256.Sum256(combined_HashInput)
	}
	// Print the final integrity hash of the entire file
	fmt.Printf("Final file integrity hash: %x\n", fullFileHash)
}

func main() {
	// Define the block size (for example, 16 bytes per block)
	const blockSize = 16

	// Perform block-by-block hashing of the file "input.txt"
	hashFileBlockByBlock_With_IntegrityCheck("input.txt", blockSize)
	fmt.Println("File hashed successfully block by block!")
}
