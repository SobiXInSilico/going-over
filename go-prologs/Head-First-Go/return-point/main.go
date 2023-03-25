package main

import (
	"fmt"
	"math"
)

func double(number float64) float64 {
	return number * 2
}

func status(grade float64) string {
	if grade < 60.00 {
		//if grade is failing return immediately
		return "failing"
	}
	//only runs if grade is >= 60
	return "passing"
}

func floatParts(number float64) (integerPart int, fractionalPart float64) {
	wholeNumber := math.Floor(number)
	return int(wholeNumber), number - wholeNumber
}

func main() {
	dozen := double(6.0)
	fmt.Println(dozen)
	fmt.Println(double(4.2))
	fmt.Println(status(60.1))
	fmt.Println(status(59))
	cans, remainder := floatParts(1.26)
	fmt.Println(cans, remainder)

}
