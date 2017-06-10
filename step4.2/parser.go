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

// ITEMREAD is any article (book, movie, ...) with the Amazon structure
type ITEMREAD struct {
	Asin  string  `json:"asin"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
	ImURL string  `json:"imUrl"`
	Related struct {
		AlsoBought []string `json:"also_bought"`
	} `json:"related"`
	Brand      string     `json:"brand"`
	Categories [][]string `json:"categories"`
}

var (
	old = []byte("'")
	new = []byte(`"`)
)

/*
PARSE ALL ITEMS & COUNT ERRORS
 */

func main() {
	var errDec error

	t0 := time.Now()
	count := 0
	countOk := 0
	countErrDec := 0
	countErrId := 0

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

		// if count%500000 == 0 {
		// 	fmt.Println(count)
		// }

		// The received file is wrong formatted (' instead of " in the json format)
		itemModif := bytes.Replace(itemRaw, old, new, -1)

		// Decode json to original struct
		errDec = json.Unmarshal(itemModif, &item)
		if errDec != nil {
			// //if item.Asin != "" {
			// fmt.Println(string(itemRaw))
			// // fmt.Println("->")
			// // fmt.Printf("%+v", item)
			// fmt.Println("")
			// // fmt.Println("---------------")
			// // }
			countErrDec++

			if countErrDec == 1 {
				fmt.Println("line", count)
				fmt.Println(string(itemRaw))
				fmt.Println(string(itemModif))
				os.Exit(0)
			}
		} else {
			countOk++
		}

		if item.Asin == "" {
			countErrId++
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
	fmt.Println(countErrId, "have an empty id")
	fmt.Println(countOk, "could be correctly read")

	// 9430088 items parsed in : 2m29.425716213s
	// 3401453 errors detected during decoding
	// 3401453 have an empty id
	// 6028635 could be correctly read
}
