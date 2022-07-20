package main

// Function to return fibonacci number
func fibonacci(n int) int {
	fb1 := 1
	fb2 := 1
	for i := 2; i < n; i++ {
		fb1, fb2 = fb2, fb1+fb2
	}
	return fb2
}
