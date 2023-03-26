package main

import "fmt"

// This code erroneously reports that the student "Carl" is failing, when in reality he just hasn’t had any grades logged.
func status(name string) {
	grades := map[string]float64{"Alma": 0, "Rohit": 86.5}
	grade := grades[name]
	if grade < 60 {
		fmt.Printf("%s is failing!\n", name)
	}
}

// to address this issue:
func status2(name string) {
	grades := map[string]float64{"Alma": 0, "Rohit": 86.5}
	grade, ok := grades[name]
	if !ok {
		fmt.Printf("No grade recorded for %s.\n", name)
	} else {
		if grade > 60 {
			fmt.Printf("%s is failing!\n", name)
		}
	}
}
func main() {
	status("Alma")
	status("Sobhan")
	counters := map[string]int{"a": 1, "b": 2}
	var value int
	var ok bool
	fmt.Println()
	value, ok = counters["a"]
	fmt.Println(value, ok)
	value, ok = counters["c"]
	fmt.Println(value, ok)
	value, ok = counters["b"]
	fmt.Println(value, ok)
	fmt.Println()
	status2("Alma")
	status2("Sobhan")
}
