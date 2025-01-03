package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Print(isPalindrome(1000021))
}

func isPalindrome(x int) bool {
	text := []rune(strconv.Itoa(x))
	fmt.Print(text)
	if len(text) <= 1 {
		return true
	}

	if text[0] != text[len(text)-1] {
		return false
	}
	result, no := strconv.Atoi(string(text[1 : len(text)-1]))
	fmt.Print(no)
	return isPalindrome(result)
}
