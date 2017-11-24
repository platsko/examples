package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

func RmDuplicates() {
	const (
		nl  = '\n'
		eof = -1
	)

	var (
		n, idx int32
		char   rune
		curr   = make([]byte, 0)
		last   []byte
		ln     = []byte("\n")
		reader = bufio.NewReader(os.Stdin)
	)

	_, _ = fmt.Scanln(&n)
	for idx = 0; idx < n; idx++ {
		for {
			char, _, _ = reader.ReadRune()
			if char == nl || char == eof {
				if last == nil {
					_, _ = os.Stdout.Write(curr)
					last = curr[:]
				} else if !reflect.DeepEqual(curr, last) {
					_, _ = os.Stdout.Write(append(ln, curr...))
					last = curr[:]
				}
				curr = nil
				break
			}
			curr = append(curr, byte(char))
		}
	}
}
