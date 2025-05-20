package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	digits := os.Args[1]
	cache := 100_000
	if len(os.Args) > 2 {
		cacheStr := os.Args[2]
		if cacheStr != "" {
			v, err := strconv.Atoi(cacheStr)
			if err == nil {
				cache = v
			}
		}
	}

	before := time.Now()

	num, err := search("pi.txt", digits, cache)
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

func search(filename, digits string, cache int) (int64, error) {
	file, err := os.Open(fmt.Sprintf("./data/%s", filename))
	if err != nil {
		return 0, err
	}
	defer file.Close()

	bins := []byte(digits)
	num := len(bins)

	var i int64
	for {
		if _, err = file.Seek(i, 0); err != nil {
			return 0, err
		}

		buf := make([]byte, cache+num)
		if _, err := file.Read(buf); err != nil {
			if err.Error() != "EOF" {
				return 0, fmt.Errorf("failed to read: %v", err)
			}

			break
		}

		got, found := searchInMemory(cache, buf, bins)
		if found {
			return i + int64(got), err
		}

		i += int64(cache)
	}

	return -1, nil
}

func searchInMemory(cache int, buf, bins []byte) (int, bool) {
outerLoop:
	for i := 0; i < cache; i++ {
		for j, bin := range bins {
			if buf[i+j] != bin {
				continue outerLoop
			}
		}

		return i, true
	}

	return 0, false
}
