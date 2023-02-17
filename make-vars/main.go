package main

import "fmt"

func main() {
	var quantity int
	var length, width float64
	var CustomerName string

	quantity = 4
	length, width = 1.2, 2.4
	CustomerName = "Sicilian Silicon"

	fmt.Println(CustomerName)
	fmt.Println("has ordered", quantity, "sheets")
	fmt.Println("each with an area of")
	fmt.Println(length*width, "square meters.")

}
