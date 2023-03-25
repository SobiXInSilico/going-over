/*
Here’s an inRange function that takes a minimum value, a maximum value,
and any number of additional float64 arguments. It will discard any
argument that is below the given minimum or above the given maximum,
returning a slice containing only the arguments that were in the specified
range.
*/
package main

import "fmt"

func inRange(min float64, max float64, numbers ...float64) []float64 {
	var result []float64
	for _, number := range numbers {
		if number >= min && number <= max {
			result = append(result, number)
		}
	}
	return result
}

func main() {
	fmt.Println(inRange(1, 100, -12.5, 3.2, 0, 50, 103.5))
	fmt.Println(inRange(-10, 10, 4.1, 12, -12, -5.2))
}
