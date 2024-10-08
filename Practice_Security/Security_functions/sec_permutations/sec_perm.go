package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	// Create a slice to store the values of permutations, with size n
	values := make([]int, n+1) // +1 to handle 1-based indexing

	// Create a slice to store the functions f(i) = Y
	invo := make([]int, n+1) // +1 to handle 1-based indexing

	// Read the function values into the slice
	for i := 1; i <= n; i++ { // Start from 1 instead of 0
		fmt.Scan(&values[i])
	}

	// Populate the invo function: invo[f(i)] = i
	for i := 1; i <= n; i++ {
		invo[values[i]] = i
	}

	for i := 1; i <= n; i++ {
		f := values[i]
		f_of_f := values[f]
		fmt.Println(f_of_f)
	}
}