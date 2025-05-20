package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	digits := os.Args[1]
	cacheSize := 100_000
	if len(os.Args) > 2 {
		cacheStr := os.Args[2]
		if cacheStr != "" {
			v, err := strconv.Atoi(cacheStr)
			if err == nil {
				cacheSize = v
			}
		}
	}

	before := time.Now()

	num, err := search("pi.txt", digits, cacheSize)
	if err != nil {
		fmt.Println(err)
	}

	after := time.Now()

	if num > 0 {
		fmt.Println(num)
	} else {
		fmt.Println("not found")
	}

	fmt.Println(after.Sub(before))
}

func search(filename, digits string, cacheSize int) (int64, error) {
	file, err := os.Open(fmt.Sprintf("./data/%s", filename))
	if err != nil {
		return 0, err
	}
	defer file.Close()

	bins := []byte(digits)
	cache := make([]byte, cacheSize+len(bins))

	for i := int64(0); ; i += int64(cacheSize) {
		if _, err = file.Seek(i, 0); err != nil {
			return 0, err
		}

		if _, err := file.Read(cache); err != nil {
			if err.Error() != "EOF" {
				return 0, fmt.Errorf("failed to read: %v", err)
			}

			break
		}

		if got, found := searchInMemory(cacheSize, cache, bins); found {
			return i + int64(got), err
		}
	}

	return -1, nil
}

func searchInMemory(cacheSize int, cache, bins []byte) (int, bool) {
outerLoop:
	for i := 0; i < cacheSize; i++ {
		for j, bin := range bins {
			if cache[i+j] != bin {
				continue outerLoop
			}
		}

		return i, true
	}

	return 0, false
}
