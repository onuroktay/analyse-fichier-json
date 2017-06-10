package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
	"sync"
	"log"

	"github.com/buger/jsonparser"
)

/* Own Data Structure */

type ITEM struct {
	ID         string   `json:"id"`
	Title      string   `json:"title"`
	Price      float64  `json:"price"`
	ImURL      string   `json:"imUrl"`
	Related    []string `json:"related"`
	Brand      string   `json:"brand"`
	Categories []string   `json:"categories"`
	//Rank   int      `json:"rank"`
}

// ---------------------------------

var (
	concurrency = 8
	sem         = make(chan bool, concurrency)
	wg          sync.WaitGroup
	database    *DATABASE
)

/*
SAVE ITEMS IN DATABASE
 */

func main() {
	t0 := time.Now()
	count := 0

	// Connect to ElasticSearch
	es, err := NewElasticSearch("amazonreader", "item")
	if err != nil {
		log.Fatalln("ElasticSearch connection error:", err.Error())
	}

	// Set Database
	database = &DATABASE{accesser: es}

	file, err := os.Open("../../json/amazon.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	r := bufio.NewReader(file)

	// Read first line
	itemRaw, err := r.ReadBytes('\n')

	for err == nil {
		count++

		// Wait as long the channel is full and then send a new message ()
		sem <- true

		// Increment the waiting group
		wg.Add(1)

		// Import next item
		go importItem(itemRaw, database)

		// Read next line
		itemRaw, err = r.ReadBytes('\n')
	}

	wg.Wait()

	// Check Error -> only EOF is a normal error
	if err != io.EOF {
		fmt.Println(err)
		return
	}

	fmt.Println(count, "items parsed and saved in :", time.Since(t0))
	// 9430088 items parsed and saved in : 1h14m58.84257534s
}

func importItem(itemRaw []byte, db *DATABASE) {
	defer wg.Done()

	item := readItem(itemRaw)

	// Save in database
	err := db.accesser.SaveItem(item)
	if err == nil {

	}

	// Read in semaphore channel to free the process
	<-sem
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
		item.Categories = append(item.Categories, string(value))
	}, "categories", "[0]")

	jsonparser.ArrayEach(raw, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		item.Related = append(item.Related, string(value))
	}, "related", "also_viewed")

	return item
}
