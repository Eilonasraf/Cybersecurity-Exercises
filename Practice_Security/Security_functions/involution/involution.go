package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	// Create a slice to store the values of f(i), with size n
	values := make([]int, n+1) // +1 to handle 1-based indexing

	// Create a slice to store the inverse function f^-1(i)
	inverse := make([]int, n+1) // +1 to handle 1-based indexing

	// Read the function values into the slice
	for i := 1; i <= n; i++ { // Start from 1 instead of 0
		fmt.Scan(&values[i])
	}

	// Populate the inverse function: inverse[f(i)] = i
	for i := 1; i <= n; i++ {
		inverse[values[i]] = i
	}

	flag := 0
	// Compare the inverse function f^-1(i) with f(i)'s values
	for i := 1; i <= n; i++ {
		if values[i] != inverse[i] {
			flag++
		}
	}

	if flag != 0 {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}
