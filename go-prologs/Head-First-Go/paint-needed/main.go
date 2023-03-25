package main

import (
	"fmt"
	"log"
)

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		// returning 0 is meaningless when there’s an error, but we had to use something for the first return value.
		return 0, fmt.Errorf("a width of %0.2f is invalid, width", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid, height", height)
	}
	area := width * height
	//The manufacturer says each liter of paint covers 10 square meters.
	return area / 10.0, nil
}

func main() {
	//err = fmt.Errorf("a height of %0.2f is invalid", -2.333333)
	//err := errors.New("height can't be begative")
	//fmt.Println(err.Error())
	//fmt.Println(err)
	//log.Fatal(err)

	var amount, total float64
	//first wall
	amount, err := paintNeeded(4.2, -3.0)
	if err != nil {
		//fmt.Println(err)
		log.Fatal(err)
	} else {
		fmt.Printf("%0.2f liters needed\n", amount)
	}

	total += amount
	//second wall
	amount, err = paintNeeded(5.2, 3.5)
	fmt.Println(err)
	fmt.Printf("%0.2f liters needed\n", amount)
	total += amount
	fmt.Printf("Total: %0.2f liters \n", total)
}
