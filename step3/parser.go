package main

import (
	"time"
	"os"
	"fmt"
	"bufio"
	"io"
)

/*
CHANGE READ METHOD - READSTRING
 */

func main() {
	t0 := time.Now()
	count := 0

	file, err := os.Open("../../json/metadata.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	r := bufio.NewReader(file)

	// Read first line
	s, err := r.ReadString('\n')

	for err == nil {
		count++

		if len(s) == 0 {
			fmt.Println("line empty")
		}

		// Read next line
		s, err = r.ReadString('\n')
	}

	// Check Error -> only EOF is a normal error
	if err != io.EOF {
		fmt.Println(err)
		return
	}

	fmt.Println(count, "items parsed in :", time.Since(t0))
	// 9430088 items parsed in : 13.6s
}