// count tallies the number of times each line
// occurs within a file
package main

import (
	"fmt"
	"github.com/SicilianSilicon/going-overn/go-prologs/Head-First-Go/datafile"
	"log"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)
}
