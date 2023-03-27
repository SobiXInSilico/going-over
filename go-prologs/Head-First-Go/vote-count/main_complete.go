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
	//candidate names as KEYs, vote counts as VALUEs
	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}
	/* for name, count := range counts {
	// 	fmt.Printf("Votes for %s : %d\n", name, count)
	} */
	//get all names in alphabetical order
	names := make([]string, 0, len(counts))
	for name := range counts {
		names = append(names, name)
	}
	// print the vote counts for each candidate in alphabetical order
	for _, name := range names {
		count := counts[name]
		fmt.Printf("Votes for %s : %d\n", name, count)
	}
}
