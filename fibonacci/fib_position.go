// Find position of Fibonacci number
package main

import "fmt"

func main() {
	a := 0
	b := 1
	count := 1
	var n, c int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		c = a + b
		a = b
		b = c
		count++
		if c == n {
			fmt.Println(count)
			break
		} else if c > n {
			fmt.Println(-1)
			break
		}
	}
}
