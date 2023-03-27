// count tallies the number of times each line
// occurs within a file
package main

import (
	"fmt"
	"log"
	"reflect"

	"github.com/SicilianSilicon/going-over/go-prologs/Head-First-Go/datafile"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\n", reflect.TypeOf(lines), ":")
	fmt.Println(lines)
}
