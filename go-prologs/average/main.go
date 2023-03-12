package main

import (
	"bufio"
	"datafile"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	//numbers := [3]float64{71.8, 56.2, 89.5}
	numbers, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)

}

func GetFloats(fileName string) ([]float64, error) {
	var numbers []float64
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	//i := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
		number, err := strconv.ParseFloat(scanner.Text(), 64)

		if err != nil {
			return nil, err
		}
		//i++
		numbers = append(numbers, number)
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return numbers, nil
}
