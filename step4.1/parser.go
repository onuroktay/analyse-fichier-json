package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

/* Original Data Structure */

type RELATED struct {
	Also_bought []string `json:"also_bought"`
}

type ITEMREAD struct {
	Asin       string     `json:"asin"`
	Title      string     `json:"title"`
	Price      float64    `json:"price"`
	ImUrl      string     `json:"imUrl"`
	Related    RELATED    `json:"related"`
	Brand      string     `json:"brand"`
	Categories [][]string `json:"categories"`
}

// ---------------------------------

func main() {
	old := []byte("'")
	new := []byte(`"`)
	t0 := time.Now()
	count := 0
	countErrDec := 0

	file, err := os.Open("../../json/metadata.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	r := bufio.NewReader(file)

	// Read first line
	itemRaw, err := r.ReadBytes('\n')

	for err == nil {
		var item ITEMREAD

		count++

		// Correction (the received file is wrong formatted: ' instead of ")
		itemRaw = bytes.Replace(itemRaw, old, new, -1)

		// Decode json to original struct
		errDec := json.Unmarshal(itemRaw, &item)
		if errDec != nil {
			fmt.Println(string(itemRaw))
			fmt.Println("exit after", count, "positions")
			fmt.Println(errDec)
			os.Exit(0)
		}

		// Read next line
		itemRaw, err = r.ReadBytes('\n')
	}

	// Check Error -> only EOF is a normal error
	if err != io.EOF {
		fmt.Println(err)
		return
	}

	fmt.Println(count, "items parsed in :", time.Since(t0))
	fmt.Println(countErrDec, "errors detected during decoding")
}
