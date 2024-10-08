# Go Cybersecurity Exercises

## Overview
This repository contains various Go-based cybersecurity exercises and projects, focusing on encryption, hashing, HTTP server handling, and security practices. The goal of this repository is to explore key concepts in Go programming and cybersecurity through hands-on examples.

## Features
- **Caesar Cipher**: Implement a basic Caesar cipher to shift letters in a string.
- **Block-by-Block Hashing with SHA-256**: Hash data block-by-block using the SHA-256 algorithm and verify data integrity.
- **Simple HTTP Server**: A Go-based HTTP server to handle static files, POST requests, and basic form handling.
- **AES Encryption/Decryption**: Securely encrypt and decrypt data using the AES algorithm.
- **Practice Security Exercises**: Various security-related exercises from HackerRank to enhance cybersecurity knowledge.

## Technologies Used
- **Programming Language**: Go
- **Encryption**: AES, SHA-256
- **HTTP Server**: Go’s built-in net/http package
- **Testing**: Go's built-in testing framework and Postman for manual API testing

## Project Structure
Each project is organized in its own folder:

- **caesar_cipher/**: Implementation of Caesar cipher encryption and decryption.
- **hashing/**: Block-by-block hashing using SHA-256, with data integrity verification.
- **http_server/**: A basic HTTP server built using Go, serving static files and handling form submissions.
- **encryption/**: AES encryption and decryption.
- **practice_security/**: Miscellaneous security exercises from HackerRank, focused on cybersecurity principles.

## Getting Started

### Prerequisites
Make sure you have the following installed:

- **Go**: 1.18 or later

### Running the Projects

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/Eilonasraf/Go-Cybersecurity-Exercises.git
   cd Go-Cybersecurity-Exercises
2. **Navigate to the desired folder and run the corresponding Go program. Example for caesar_cipher:**:
   ```bash
   cd caesar_cipher
   go run main.go

## Running Tests

1. **Some projects include test cases that can be run using Go’s built-in testing framework. Example for hash_sha256_blockbyblock:**:
   ```bash
   cd hashing/hash_sha256_blockbyblock
   go test

**This project is licensed under the MIT License.**
