// pass_fail reports wether a grade is passing or failing
package main

import (
	"fmt"
	"github.com/SicilianSilicon/going-over/go-prologs/keyboard"
	"log"
)

func main() {
	fmt.Print("Enter a grade: ")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "faiing"
	}

	fmt.Println("A grade of", grade, "is", status)
}
