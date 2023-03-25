package main

import "fmt"

func main() {

	notes := [7]string{"do", "re", "me", "fa", "so", "la", "si"}
	for i := 0; i < len(notes); i++ {
		fmt.Println(i+1, notes[i])
	}

	fmt.Println("")

	//for...range form has no messy init, condition, and post expressions.
	for index, note := range notes {
		fmt.Println(index+1, note)
	}

	fmt.Println("")
	// this time with blank identifier
	for _, note := range notes {
		fmt.Println(note)

	}

}
