package main

import (
	"fmt"
	"unicode"
)

func main() {
	input := "Madam,I'mAdam"
	normalized := ""
	for _, r := range input {
		if unicode.IsLetter(r) {
			normalized += string(unicode.ToLower(r))
		}
	}
	isPalindrome := true
	for i := 0; i < len(normalized)/2; i++ {
		if normalized[i] != normalized[len(normalized)-1-i] {
			isPalindrome = false
			break
		}
	}
	if isPalindrome {
		fmt.Println("It's a palindrome!")
	} else {
		fmt.Println("It's not a palindrome.")
	}
}
