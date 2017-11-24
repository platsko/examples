package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"unicode"
)

func Merge() {
	var (
		in, idx int
		list    []int
	)

	_, _ = fmt.Scanln(&in)
	reader := bufio.NewReader(os.Stdin)
	for idx = 0; idx < in; idx++ {
		num := readItem(reader)
		for l := num; num >= 0 && l > 0; l-- {
			if num = readItem(reader); num >= 0 {
				list = append(list, num)
			}
		}
	}

	buff := make([]byte, 0)
	if list != nil {
		sort.Ints(list)
		s := []byte(" ")
		ll := len(list)
		last := ll - 1
		for idx = 0; idx < ll; idx++ {
			buff = append(buff, []byte(strconv.Itoa(list[idx]))...)
			if idx < last {
				buff = append(buff, s...)
			}
		}
	}

	_, _ = os.Stdout.Write(buff)
}

func readItem(reader *bufio.Reader) int {
	buff := make([]rune, 0)
	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		if r == ' ' {
			if len(buff) == 0 {
				continue
			}
			num, err := strconv.Atoi(string(buff))
			buff = make([]rune, 0)
			if err != nil {
				continue
			}
			return num
		} else if unicode.IsDigit(r) {
			buff = append(buff, r)
		}
	}

	return -1
}
