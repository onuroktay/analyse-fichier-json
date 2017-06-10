package main

import (
	"os"
	"fmt"
	"bufio"
)

/*
CHECK FILE CONTENT
 */
func main() {
	count := 0

	file, err := os.Open("../../json/metadata.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)

	for scan.Scan() {
		fmt.Println(scan.Text())
		count++

		// Read 3 first items and stop
		if count == 3 {
			return
		}
	}
}
