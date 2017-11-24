package main

import (
	"fmt"
	"strings"
)

func IsPalindrome(source string) string {
	// cast to lower case and remove spaces
	s := strings.ReplaceAll(strings.ToLower(source), " ", "")

	l := len(s)
	n := l - 1
	c := make([]rune, l, l)
	for _, char := range s {
		c[n] = char
		n--
	}

	if s == string(c[n+1:]) { // cut lead spaces and compare
		return fmt.Sprintf("%s - is palindrome.", source)
	}

	return fmt.Sprintf("%s - is not palindrome.", source)
}
