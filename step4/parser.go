package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

/* Original Data Structure */

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

// ---------------------------------

/*
PARSE 2nd ITEM FOR CHECK
 */
func main() {
	old := []byte("'")
	new := []byte(`"`)
	count := 0

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

		// Stop after second item
		if count == 2 {
			// The received file is wrong formatted (' instead of " in the json format)
			itemRaw = bytes.Replace(itemRaw, old, new, -1)

			// Cast read item string to []byte, then unmarshaling result into struct
			json.Unmarshal(itemRaw, &item)
			fmt.Println(string(itemRaw))
			fmt.Println("->")
			fmt.Printf("%+v\n", item)

			return
		}

		// Read next line
		itemRaw, err = r.ReadBytes('\n')
	}

	// Check Error -> only EOF is a normal error
	if err != io.EOF {
		fmt.Println(err)
		return
	}
}
