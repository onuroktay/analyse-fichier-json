package main

import (
	"bufio"
	"log"
	"io"
	"os"
	"time"
	"fmt"
)

var (
	counter   = 0
	mark1     = `'`
	mark2     = `"`
	backslash = `\`
	t0        = time.Now()
)

/*
JSON CORRECTOR -> GENERATE NEW FILE (amazon.json)
 */

func main() {
	s := ""

	// Open file to read
	fileR, err := os.Open("../../json/metadata.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer fileR.Close()

	// Open file for writing
	fileW, err := os.Create("../../json/amazon.json")
	if err != nil {
		log.Fatal(err)
	}
	defer fileW.Close()

	// Create a buffered reader from the file
	r := bufio.NewReader(fileR)

	// Create a buffered writer from the file
	w := bufio.NewWriter(fileW)

	// Read first line
	s, err = r.ReadString('\n')

	for err == nil {
		counter++

		// Write in buffer
		w.WriteString(CorrectQuotes(s))

		// Read next line
		s, err = r.ReadString('\n')
	}

	// Write memory buffer to disk
	w.Flush()

	// Check Error -> only EOF is a normal error
	if err != io.EOF {
		log.Println(err)
		return
	}

	log.Println(counter, "items parsed succesfully in :", time.Since(t0))
	// 9430088 items parsed succesfully in : 41m23.702319783s
	// -> 0.88us / operation, or, compared to the 10.5 GB, 1.27 GB data processed by second
}

func CorrectQuotes(s1 string) (s2 string) {
	inside := false
	marker := ""
	c := ""

	// Display counter every 100'000 items
	if (counter % 100000) == 0 {
		fmt.Println(counter, "items saved in :", time.Since(t0))
	}

	for pos, char := range s1 {
		c = string(char)

		// Detect Start Marker
		if !inside && (c == mark1 || c == mark2) {
			inside = true
			s2 += mark2

			if c == mark1 {
				marker = mark1
			} else {
				marker = mark2
			}

			continue
		}

		// Detect End Marker
		if inside {
			if c == marker && s1[pos-1:pos] != backslash {
				inside = false
				s2 += mark2

				continue
			}
		}

		// Replace " by '
		if inside && (c == mark2) {
			s2 += mark1
			continue
		}

		// Ignore \
		if inside && (c == backslash) {
			continue
		}

		// Transfer char
		s2 += c
	}

	return
}
