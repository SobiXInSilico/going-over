// pass_fail reports wether a grade is passing or failing
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a grade: ")
	//reader :=
	//will not compile
	//input := reader.ReadString('\n')
	//will not compile. err declared and not used
	//input, err := reader.ReadString('\n')
	//option 1 : use blank identifier
	//input, _ := reader.ReadString('\n')
	//option 2 : handle the error _ but every time we runt it , we see the nil value
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	//log.Fatal(err)
	//so, we use if to see the err only when it has something to offer
	if err != nil {
		log.Fatal(err)
	}
	/* now we passed the error handling. let's evaluate the grade.
	here’s the new problems:
		1)input from the keyboard is read in as a string. strconv.ParseFloat()
		2)the input string still has a newline character on the end, from when the user pressed the Enter key while entering it. we need to strip that off. strings.TrimSpace()*/

	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	//make sure that the scope has the status
	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "faiing"
	}

	fmt.Println("A grade of", grade, "is", status)
}
