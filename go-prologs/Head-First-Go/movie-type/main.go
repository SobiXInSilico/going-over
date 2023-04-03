package main

import "fmt"

type movie struct {
	director string
	year     int
	made     string
}

func showInfo(p movie) {
	fmt.Println("Director:", p.director)
	fmt.Println("Year:", p.year)
	fmt.Println("Made in:", p.made)
}
func main() {
	var breathless movie
	breathless.director = "Jean-Luc Godard"
	breathless.year = 1959
	breathless.made = "France"
	showInfo(breathless)
}
