// count tallies the number of times each line
// occurs within a file
package main

import (
	"fmt"
	"log"

	"github.com/SicilianSilicon/going-over/go-prologs/Head-First-Go/datafile"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}
	fmt.Println(counts)
}
