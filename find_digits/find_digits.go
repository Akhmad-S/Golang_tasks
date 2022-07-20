// Two numbers are given. Determine the digits included in the record of both the first and second numbers.
// But provided that the digits in the numbers do not repeat
package main

import "fmt"

func main() {
	var firstNum, secNum, revFirstNum int
	fmt.Scan(&firstNum, &secNum)

	for 0 < firstNum {
		revFirstNum *= 10
		revFirstNum += firstNum % 10
		firstNum /= 10
	}
	for i := revFirstNum; 0 < i; i /= 10 {
		n1 := i % 10
		for j := secNum; 0 < j; j /= 10 {
			n2 := j % 10
			if n1 == n2 {
				fmt.Print(n1, " ")
			}
		}
	}
}
