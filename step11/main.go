package main

import (
	"flag"
	"fmt"
	"github.com/onuroktay/amazon-reader/Analyse_fichier_json/step10"
)

func main() {
	fileName := flag.String("file", "amazon.json", "name of the json file to read")
	flag.Parse()

	fmt.Println("Import data from ", *fileName)
	err := OnurTPIjsonReader.ImportJSON(*fileName)
	if err != nil {
		fmt.Println(err.Error())
	}
}
