// AsciiCounter scans file list in dir and build
// histogram dictionary with counter of ASCII symbols.

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
)

func main() {
	const dir = "path/to/dir"
	var (
		// dict contains counters by byte
		dict = make([]uint64, 128, 128)
		// can't print bytes names mapping
		bToStr = map[byte]string{
			0:   "NUL",
			7:   "BEL",
			8:   "BS",
			9:   "TAB",
			10:  "LF",
			13:  "CR",
			32:  "SP",
			127: "DEL",
		}
		// wait group before print
		wg = &sync.WaitGroup{}
	)

	// walks the dir tree rooted
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.Size() > 0 { // only not empty files
			wg.Add(1) // add task to wait group
			go readAndCount(path, dict, wg)
		}
		return nil
	})
	wg.Wait() // wait for all tasks done

	if err != nil {
		fmt.Println(err)
	} else {
		printResult(dict, bToStr)
	}
}

func readAndCount(path string, dict []uint64, wg *sync.WaitGroup) {
	defer wg.Done() // done task before return
	file, err := os.Open(path)
	if err != nil {
		return // skip files can't be opened
	}
	if blob, err := ioutil.ReadAll(file); err == nil {
		go file.Close()
		for _, b := range blob {
			if b < 128 { // the byte belongs to ASCII table
				atomic.AddUint64(&dict[b], 1)
			}
		}
	}
}

func printResult(dict []uint64, bToStr map[byte]string) {
	l := len(dict)     // length
	total := uint64(0) // total counter
	for i := 0; i < l; i++ {
		if c := dict[i]; c > 0 {
			b := byte(i)
			total += c
			if n, ok := bToStr[b]; ok {
				fmt.Printf("'%s': %d\n", n, c)
			} else {
				fmt.Printf("'%s': %d\n", string(b), c)
			}
		}
	}

	// print total counter
	fmt.Printf("Total: %d", total)
}
