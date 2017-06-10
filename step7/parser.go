package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
	"github.com/buger/jsonparser"
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
PARSE ALL ITEM WITH JSONPARSER INSTEAD OF JSON.UNMARSHAL
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

		readItem(itemRaw)

		// Read next line
		itemRaw, err = r.ReadBytes('\n')
	}

	// Check Error -> only EOF is a normal error
	if err != io.EOF {
		fmt.Println(err)
		return
	}

	fmt.Println(count, "items parsed in :", time.Since(t0))
	// 9430088 items parsed in : 1m27.562249871s
}

func readItem(raw []byte) *ITEM {
	item := &ITEM{}

	// Read properties from json
	item.ID, _ = jsonparser.GetString(raw, "asin")
	item.Title, _ = jsonparser.GetString(raw, "title")
	item.Price, _ = jsonparser.GetFloat(raw, "price")
	item.ImURL, _ = jsonparser.GetString(raw, "imUrl")
	item.Brand, _ = jsonparser.GetString(raw, "brand")

	jsonparser.ArrayEach(raw, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		item.Category = string(value)
	}, "categories", "[0]")

	jsonparser.ArrayEach(raw, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		item.Related = append(item.Related, string(value))
	}, "related", "also_viewed")

	//jsonparser.ObjectEach(itemRaw, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
	//	item.Rank = string(value)
	//	return nil
	//}, "salesRank")

	return item
}
