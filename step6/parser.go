package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
	"encoding/json"
)

/* Own Data Structure */

type ITEM struct {
	ID       string   `json:"id"`
	Title    string   `json:"title"`
	Price    float64  `json:"price"`
	ImURL    string   `json:"imUrl"`
	Related  []string `json:"related"`
	Brand    string   `json:"brand"`
	Category string   `json:"category"`
	//Rank   int      `json:"rank"`
}

// ---------------------------------

/*
PARSE ALL ITEMS
 */

func main() {
	t0 := time.Now()
	count := 0

	file, err := os.Open("../../json/amazon.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	r := bufio.NewReader(file)

	// Read first line
	itemRaw, err := r.ReadBytes('\n')
	if itemRaw == nil {

	}

	for err == nil {
		count++

		item := &ITEM{}
		json.Unmarshal(itemRaw, &item)

		// Read next line
		itemRaw, err = r.ReadBytes('\n')
	}

	// Check Error -> only EOF is a normal error
	if err != io.EOF {
		fmt.Println(err)
		return
	}

	fmt.Println(count, "items parsed in :", time.Since(t0))
	// 9430088 items parsed in : 2m7.973722313s
}
