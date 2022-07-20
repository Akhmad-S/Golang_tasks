package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	palindrome(str)
}

func palindrome(text string) {
	var newText string
	text = strings.Trim(text, "\n")
	text = strings.ReplaceAll(text, " ", "")
	text = strings.ToLower(text)
	r := []rune(text)

	for i := len(r) - 1; i >= 0; i-- {
		newText += string(r[i])
	}
	if newText == text {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not palindrome")
	}
}
