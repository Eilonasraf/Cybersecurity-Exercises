package main

import "fmt"

func main() {
	// Variable to store the size of the set
	var n int
	// Read the input size n from standard input
	fmt.Scan(&n)

	// Create a slice to store the values of f(i), with size n
	values := make([]int, n)

	// Create a map (similar to a dictionary in Python),
	// to track whether a value has been seen before.
	// The keys in this map represent the values from f(i),
	// and the boolean represents whether the value was encountered.
	found := make(map[int]bool)

	// Flag to check if the function is bijective (onto and one-to-one) (starts as true,
	// turns false if any condition fails)
	isBijective := true

	// Read the values into the slice
	for i := 0; i < n; i++ {
		fmt.Scan(&values[i])

		// Check if the value is outside the valid range [1, n]
		// A value outside this range means the function is not surjective (onto)
		if values[i] < 1 || values[i] > n {
			isBijective = false
		}

		// Check if the value has been seen before (to ensure injective (one-to-one) property,
		// i.e., no duplicates)
		// If we have seen the value before, the function is not injective (one-to-one)
		if found[values[i]] {
			isBijective = false
		}
		// Mark this value as found in the map
		found[values[i]] = true
	}

	// Output "YES" if the function is bijective (both injective and surjective)
	// Otherwise, output "NO"
	if isBijective {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}