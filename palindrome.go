package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	var input string
	fmt.Println("Enter Your String To Check Palindrome:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	input = strings.TrimSpace(input)
	normalized := ""
	for _, r := range input {
		if unicode.IsLetter(r) {
			normalized += string(unicode.ToLower(r))
		}
	}
	palindrome := true
	length := len(normalized)
	for i := 0; i < length/2; i++ {
		if normalized[i] != normalized[length-1-i] {
			palindrome = false
			break
		}
	}
	if palindrome {
		fmt.Println("Your Given String Is Palindrome.")
	} else {
		fmt.Println("Your Given String Is Not Palindrome!")
	}
}
