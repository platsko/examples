package main

import (
	"fmt"
)

func CombineBraces(s string, opened, closed, n int) {
	if len(s) == 2*n {
		println(s)
		return
	}

	if opened < n { // add next open-brace
		CombineBraces(s+"(", opened+1, closed, n)
	}

	if closed < opened { // add next close-brace
		CombineBraces(s+")", opened, closed+1, n)
	}
}

func ParseBraces(s string) string {
	stack := make([]rune, 0)
	for _, char := range []rune(s) {
		switch char {
		case '(':
			stack = append(stack, char)
		case ')':
			l := len(stack)
			if l == 1 {
				stack = make([]rune, 0)
			} else if l > 1 {
				stack = stack[:l-1]
			} else {
				return fmt.Sprint(s, " => false")
			}
		default:
			return fmt.Sprintf("unsupported char: %s", string(char))
		}
	}

	if len(stack) == 0 {
		return fmt.Sprintf("%s => true", s)
	}

	return fmt.Sprintf("%s => false", s)
}
