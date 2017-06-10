package main

import (
	"time"
	"os"
	"fmt"
	"bufio"
)

/*
COUNT ITEMS
 */

func main() {
	t0 := time.Now()
	count := 0

	file, err := os.Open("../../json/metadata.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)

	for scan.Scan() {
		count++
	}

	fmt.Println(count, "items parsed in :", time.Since(t0))
	//245442 items parsed in : 489.786784ms
	// !!!!! -> what about 9,4 Mio items ?
}
