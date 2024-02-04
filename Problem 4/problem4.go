package main

import "fmt"

func sum_to_n_a(n int) int {
	// O(1) time complexity
	// O(1) space complexity
	result := int(((1 + n) * (n)) / 2)
	return result
}

func sum_to_n_b(n int) int {
	// O(n) time complexity
	// O(1) space complexity
	result := 0
	for i := 1; i <= n; i++ {
		result += i
	}
	return result
}

func sum_to_n_c(n int) int {
	// O(n) time complexity
	// O(n) space complexity
	if n == 1 {
		return 1
	}
	result := n + sum_to_n_c(n-1)
	return result
}

func main() {
	n := 10
	sum := sum_to_n_a(n)
	fmt.Printf("Sum of integers from 1 to %d: %.d\n", n, sum)

	sum2 := sum_to_n_b(n)
	fmt.Printf("Sum of integers from 1 to %d: %.d\n", n, sum2)

	sum3 := sum_to_n_c(n)
	fmt.Printf("Sum of integers from 1 to %d: %.d\n", n, sum3)
}
